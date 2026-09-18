package web_transport_http

import (
	"net/http"
	core_logger "neurox/internal/core/logger"
	core_http_response "neurox/internal/core/transport/http/response"
)

func (h *WebHTTPHandler) GetTermsPage(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	html, err := h.webService.GetTermsPage()
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get terms.html for main page",
		)
		return
	}

	responseHandler.HTMLResponse(html)
}
