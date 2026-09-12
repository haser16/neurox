package core_http_server

import (
	"net/http"
	core_middleware "neurox/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []core_middleware.Middleware
}

func (r *Route) WithMiddleware() http.Handler {
	return core_middleware.ChainMiddleware(
		r.Handler,
		r.Middleware...,
	)
}
