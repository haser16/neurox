package users_transport_http

import core_domain "neurox/internal/core/domain"

type UserRequest struct {
	Username string  `json:"username" validate:"required,min=2,max=100"`
	Email    string  `json:"email" validate:"required,email"`
	Phone    *string `json:"phone" validate:"omitempty,startswith=+"`
	Password string  `json:"password" validate:"required,min=8,max=30"`
}

type UserResponse struct {
	ID       int64   `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Phone    *string `json:"phone"`
	Avatar   *string `json:"avatar"`
}

func dtoFromDomain(domain core_domain.User) UserResponse {
	return UserResponse{
		ID:       domain.ID,
		Username: domain.Username,
		Email:    domain.Email,
		Phone:    domain.Phone,
		Avatar:   domain.Avatar,
	}
}
