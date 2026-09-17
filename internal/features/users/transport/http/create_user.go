package users_transport_http

import (
	"net/http"
	core_domain "neurox/internal/core/domain"
	core_logger "neurox/internal/core/logger"
	core_http_request "neurox/internal/core/transport/http/requests"
	core_http_response "neurox/internal/core/transport/http/response"
)

type CreateUserRequest UserRequest

type CreateUserResponse UserResponse

// CreateUser godoc
// @Summary Create new user
// @Description Create new user in system
// @Tags users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Create user request body"
// @Success 201 {object} CreateUserResponse "Successfully created user"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users [post]
func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate requests",
		)
		return
	}

	userDomain := domainFromDTO(request)

	var err error
	userDomain, err = h.usersService.CreateUser(ctx, userDomain)

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create user",
		)
		return
	}

	response := CreateUserResponse(dtoFromDomain(userDomain))
	responseHandler.JsonResponse(response, http.StatusCreated)
}

func domainFromDTO(request CreateUserRequest) core_domain.User {
	return core_domain.NewUserUninitialized(
		request.Username,
		request.Email,
		request.Phone,
		request.Password)
}
