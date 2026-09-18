package users_transport_http

import (
	"net/http"
	core_errors "neurox/internal/core/errors"
	core_logger "neurox/internal/core/logger"
	core_middleware "neurox/internal/core/transport/http/middleware"
	core_http_response "neurox/internal/core/transport/http/response"
)

type GetUserByJWTResponse UserResponse

// GetUserByJWT godoc
// @Summary Get user by jwt
// @Description Get user by jwt
// @Tags users
// @Produce json
// @Success 200 {object} GetUserByJWTResponse "Successfully retrieved user"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "User not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users [get]
func (h *UsersHTTPHandler) GetUserByJWT(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	userID, ok := ctx.Value(core_middleware.UserIDContextKey).(int64)
	if !ok {
		responseHandler.ErrorResponse(
			core_errors.ErrInvalidArgument,
			"UserID is missing in context",
		)
		return
	}
	userDomain, err := h.usersService.GetUserByID(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user id",
		)
		return
	}
	response := GetUserByJWTResponse(dtoFromDomain(userDomain))
	responseHandler.JsonResponse(response, http.StatusOK)
}
