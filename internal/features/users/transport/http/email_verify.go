package users_transport_http

import (
	"net/http"
	core_logger "neurox/internal/core/logger"
	core_http_request "neurox/internal/core/transport/http/requests"
	core_http_response "neurox/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) EmailVerify(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	token, err := core_http_request.GetStringQueryParam(r, "token")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get token",
		)
		return
	}

	if err := h.usersService.EmailVerify(ctx, token); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to verify email",
		)
		return
	}

	rw.Header().Set("Content-Type", "text/html; charset=utf-8")

	http.ServeFile(
		rw,
		r,
		"public/email-verify.html",
	)
}
