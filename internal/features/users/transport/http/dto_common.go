package users_transport_http

import core_domain "neurox/internal/core/domain"

type UserRequest struct {
	Username string  `json:"username" validate:"required,min=2,max=100" example:"ivanich123"`
	Email    string  `json:"email" validate:"required,email" example:"ivanich123@gmail.com"`
	Phone    *string `json:"phone" validate:"omitempty,startswith=+" example:"+375296547812"`
	Password string  `json:"password" validate:"required,min=8,max=30" example:"password123@@"`
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
