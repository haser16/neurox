package web_transport_http

import (
	core_http_server "neurox/internal/core/transport/http/server"
)

type WebHTTPHandler struct {
	webService WebService
}

type WebService interface {
	GetMainPage() ([]byte, error)
	GetLoginPage() ([]byte, error)
	GetRegisterPage() ([]byte, error)
	GetPrivacyPage() ([]byte, error)
	GetTermsPage() ([]byte, error)
	GetSupportPage() ([]byte, error)
	GetForgotPasswordPage() ([]byte, error)
	GetProfilePage() ([]byte, error)
	GetEditProfilePage() ([]byte, error)
	GetRequestsHistoryPage() ([]byte, error)
	GetCreateRequestPage() ([]byte, error)
}

func NewWebHTTPHandler(webService WebService) *WebHTTPHandler {
	return &WebHTTPHandler{
		webService: webService,
	}
}

func (h *WebHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Path:    "/",
			Handler: h.GetMainPage,
		},
		{
			Path:    "/login",
			Handler: h.GetLoginPage,
		},
		{
			Path:    "/register",
			Handler: h.GetRegisterPage,
		},
		{
			Path:    "/privacy",
			Handler: h.GetPrivacyPage,
		},
		{
			Path:    "/terms",
			Handler: h.GetTermsPage,
		},
		{
			Path:    "/support",
			Handler: h.GetSupportPage,
		},
		{
			Path:    "/forgot-password",
			Handler: h.GetForgotPasswordPage,
		},
		{
			Path:    "/profile",
			Handler: h.GetProfilePage,
		},
		{
			Path:    "/profile/edit",
			Handler: h.GetEditProfilePage,
		},
		{
			Path:    "/requests",
			Handler: h.GetRequestsHistoryPage,
		},
		{
			Path:    "/requests/create",
			Handler: h.GetCreateRequestPage,
		},
	}
}
