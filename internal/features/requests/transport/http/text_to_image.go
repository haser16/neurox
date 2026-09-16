package request_transport_http

import (
	"net/http"
	core_domain "neurox/internal/core/domain"
	core_errors "neurox/internal/core/errors"
	core_logger "neurox/internal/core/logger"
	core_middleware "neurox/internal/core/transport/http/middleware"
	core_http_request "neurox/internal/core/transport/http/requests"
	core_http_response "neurox/internal/core/transport/http/response"
)

type TextToImageRequest struct {
	Prompt string `json:"prompt"`
}

type TextToImageResponse struct {
	ID     int64  `json:"id"`
	Image  string `json:"image"`
	Prompt string `json:"prompt"`
}

func (h *RequestsHTTPHandler) TextToImage(rw http.ResponseWriter, r *http.Request) {
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

	var request TextToImageRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate requests",
		)
		return
	}

	requestDomain := domainFromDTO(request, userID)

	requestDomain, err := h.requestsService.TextToImage(ctx, requestDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to fetch text to image",
		)
	}

	response := dtoFromDomain(requestDomain)
	responseHandler.JsonResponse(response, http.StatusCreated)
}

func domainFromDTO(request TextToImageRequest, userID int64) core_domain.Request {
	return core_domain.NewRequestUninitialized(request.Prompt, userID)
}

func dtoFromDomain(requestDomain core_domain.Request) TextToImageResponse {
	return TextToImageResponse{
		requestDomain.ID,
		requestDomain.Image,
		requestDomain.Prompt,
	}
}
