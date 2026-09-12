package users_transport_http

import (
	"net/http"
	core_logger "neurox/internal/core/logger"
	core_http_request "neurox/internal/core/transport/http/requests"
	core_http_response "neurox/internal/core/transport/http/response"
	users_service "neurox/internal/features/users/service"
)

func (h *UsersHTTPHandler) UploadAvatar(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get `UserID` path value",
		)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to parse multipart form",
		)
		return
	}

	file, header, err := r.FormFile("avatar")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to parse file",
		)
		return
	}
	defer file.Close()

	avatar := users_service.Avatar{
		File:        file,
		ContentType: header.Header.Get("Content-Type"),
		Size:        header.Size,
	}

	if err := h.usersService.UploadAvatar(ctx, userID, avatar); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to upload avatar",
		)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
