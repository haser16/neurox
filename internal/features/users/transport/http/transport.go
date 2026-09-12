package users_transport_http

import (
	"context"
	"net/http"
	core_domain "neurox/internal/core/domain"
	core_http_server "neurox/internal/core/transport/http/server"
	users_service "neurox/internal/features/users/service"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		userDomain core_domain.User,
	) (core_domain.User, error)
	GetUserByID(
		ctx context.Context,
		userID int64,
	) (core_domain.User, error)
	GetUserByUsername(
		ctx context.Context,
		username string,
	) (core_domain.User, error)
	AuthenticateUser(
		ctx context.Context,
		username string,
		password string,
	) (string, error)
	DeleteUser(
		ctx context.Context,
		userID int64,
	) error
	UploadAvatar(
		ctx context.Context,
		userID int64,
		avatar users_service.Avatar,
	) error
}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUserByID,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUserByUsername,
		},
		{
			Method:  http.MethodPost,
			Path:    "/users/authenticate",
			Handler: h.AuthenticateUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/users/{id}/avatar",
			Handler: h.UploadAvatar,
		},
	}
}
