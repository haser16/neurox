package users_transport_http

import (
	"net/http"
	core_logger "neurox/internal/core/logger"
	core_http_request "neurox/internal/core/transport/http/requests"
	core_http_response "neurox/internal/core/transport/http/response"
)

type AuthenticateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticateUserResponse struct {
	Token string `json:"token"`
}

func (h *UsersHTTPHandler) AuthenticateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request AuthenticateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate request",
		)
		return
	}

	jwtToken, err := h.usersService.AuthenticateUser(ctx, request.Username, request.Password)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to authenticate user",
		)
		return
	}
	response := AuthenticateUserResponse{
		Token: jwtToken,
	}
	responseHandler.JsonResponse(response, http.StatusOK)
}
