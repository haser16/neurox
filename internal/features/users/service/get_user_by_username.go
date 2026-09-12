package users_service

import (
	"context"
	"fmt"
	core_domain "neurox/internal/core/domain"
)

func (s *UsersService) GetUserByUsername(
	ctx context.Context,
	email string,
) (core_domain.User, error) {
	userDomain, err := s.usersRepository.GetUserByUsername(ctx, email)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("failed to get user by username: %w", err)
	}
	return userDomain, nil
}
