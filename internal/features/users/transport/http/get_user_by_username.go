package users_transport_http

import (
	"net/http"
	core_logger "neurox/internal/core/logger"
	core_http_request "neurox/internal/core/transport/http/requests"
	core_http_response "neurox/internal/core/transport/http/response"
)

type GetUserByEmailResponse UserResponse

func (h *UsersHTTPHandler) GetUserByUsername(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	username, err := core_http_request.GetStringQueryParam(r, "username")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get email query param",
		)
		return
	}

	userDomain, err := h.usersService.GetUserByUsername(ctx, username)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user by username",
		)
		return
	}

	response := GetUserByEmailResponse(dtoFromDomain(userDomain))
	responseHandler.JsonResponse(response, http.StatusOK)
}
