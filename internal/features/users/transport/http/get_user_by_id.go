package users_transport_http

import (
	"net/http"
	core_logger "neurox/internal/core/logger"
	core_http_request "neurox/internal/core/transport/http/requests"
	core_http_response "neurox/internal/core/transport/http/response"
)

type GetUserByIDResponse UserResponse

func (h *UsersHTTPHandler) GetUserByID(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get id path param",
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
	response := GetUserByIDResponse(dtoFromDomain(userDomain))
	responseHandler.JsonResponse(response, http.StatusOK)
}
