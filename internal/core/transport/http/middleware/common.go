package core_middleware

import (
	"context"
	"fmt"
	"net/http"
	auth_jwt "neurox/internal/auth/jwt"
	core_errors "neurox/internal/core/errors"
	core_logger "neurox/internal/core/logger"
	core_http_response "neurox/internal/core/transport/http/response"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	RequestIDHeader     = "X-Request-Id"
	AuthorizationHeader = "Authorization"
	UserIDContextKey    = "userID"
)

type Claims struct {
	UserID int64 `json:"id"`
	jwt.RegisteredClaims
}

func CORS() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowedOrigins := map[string]struct{}{
				"http://localhost:5050": {},
				"http://127.0.0.1:5050": {},
			}
			origin := r.Header.Get("Origin")
			if _, ok := allowedOrigins[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			}
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(RequestIDHeader)

			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(RequestIDHeader, requestID)
			w.Header().Set(RequestIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(RequestIDHeader)

			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)

			ctx := core_logger.ToContext(r.Context(), l)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Auth() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			authHeader := r.Header.Get(AuthorizationHeader)

			if authHeader == "" {
				responseHandler.ErrorResponse(
					core_errors.ErrInvalidArgument,
					"Missing authorization header",
				)
				return
			}

			parts := strings.Fields(authHeader)

			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				responseHandler.ErrorResponse(
					core_errors.ErrInvalidArgument,
					"Invalid authorization header",
				)
				return
			}

			tokenString := parts[1]

			claims := &Claims{}

			token, err := jwt.ParseWithClaims(
				tokenString,
				claims,
				func(token *jwt.Token) (any, error) {
					if token.Method != jwt.SigningMethodHS256 {
						return nil, fmt.Errorf(
							"unexpected signing method: %s",
							token.Method.Alg(),
						)
					}

					return []byte(auth_jwt.NewConfigMust().Secret), nil
				},
			)

			if err != nil {
				log.Error(
					"JWT validation failed",
					zap.Error(err),
				)

				responseHandler.ErrorResponse(
					core_errors.ErrInvalidArgument,
					"Invalid token",
				)
				return
			}

			if !token.Valid {
				responseHandler.ErrorResponse(
					core_errors.ErrInvalidArgument,
					"Invalid token",
				)
				return
			}

			if claims.UserID == 0 {
				responseHandler.ErrorResponse(
					core_errors.ErrInvalidArgument,
					"UserID is missing",
				)
				return
			}

			ctx = context.WithValue(
				ctx,
				UserIDContextKey,
				claims.UserID,
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(p, "during handle HTTP requests got unexpected panic")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_response.NewResponseWriter(w)

			before := time.Now()
			log.Debug(
				">>> incoming HTTP requests",
				zap.String("http_method", r.Method),
				zap.Time("time", before.UTC()),
			)

			next.ServeHTTP(rw, r)

			log.Debug("<<< done HTTP requests",
				zap.Int("status_code", rw.GetStatusCode()),
				zap.Duration("latency", time.Now().Sub(before)))
		})
	}
}
