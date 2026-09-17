package request_transport_http

import (
	"fmt"
	"net/http"
	core_domain "neurox/internal/core/domain"
	core_errors "neurox/internal/core/errors"
	core_logger "neurox/internal/core/logger"
	core_middleware "neurox/internal/core/transport/http/middleware"
	core_http_request "neurox/internal/core/transport/http/requests"
	core_http_response "neurox/internal/core/transport/http/response"
)

type GetRequestsResponse struct {
	ID     int64  `json:"id"`
	Image  string `json:"image"`
	Prompt string `json:"prompt"`
}

func (h *RequestsHTTPHandler) GetRequests(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get limit and offset query params",
		)
		return
	}

	userID, ok := ctx.Value(core_middleware.UserIDContextKey).(int64)
	if !ok {
		responseHandler.ErrorResponse(
			core_errors.ErrInvalidArgument,
			"UserID is missing in context",
		)
		return
	}

	requestsDomain, err := h.requestsService.GetRequests(ctx, limit, offset, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get requests",
		)
		return
	}
	response := dtosFromDomainsGetRequest(requestsDomain)
	responseHandler.JsonResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)
	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' querry param: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' querry param: %w", err)
	}

	return limit, offset, nil
}

func dtoFromDomainGetRequest(requests core_domain.Request) GetRequestsResponse {
	return GetRequestsResponse{
		ID:     requests.ID,
		Image:  requests.Image,
		Prompt: requests.Prompt,
	}
}

func dtosFromDomainsGetRequest(requests []core_domain.Request) []GetRequestsResponse {
	results := make([]GetRequestsResponse, len(requests))
	for i, request := range requests {
		results[i] = dtoFromDomainGetRequest(request)
	}
	return results
}
