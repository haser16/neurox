package request_transport_http

import (
	"context"
	"net/http"
	core_domain "neurox/internal/core/domain"
	core_middleware "neurox/internal/core/transport/http/middleware"
	core_http_server "neurox/internal/core/transport/http/server"
)

type RequestsHTTPHandler struct {
	requestsService RequestsService
}

type RequestsService interface {
	TextToImage(
		ctx context.Context,
		request core_domain.Request,
	) (core_domain.Request, error)
	GetRequests(
		ctx context.Context,
		limit *int,
		offset *int,
		userID int64,
	) ([]core_domain.Request, error)
}

func NewRequestsHTTPHandler(requestsService RequestsService) *RequestsHTTPHandler {
	return &RequestsHTTPHandler{
		requestsService: requestsService,
	}
}

func (h *RequestsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/requests/text-to-image",
			Handler: h.TextToImage,
			Middleware: []core_middleware.Middleware{
				core_middleware.Auth(),
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/requests",
			Handler: h.GetRequests,
			Middleware: []core_middleware.Middleware{
				core_middleware.Auth(),
			},
		},
	}
}
