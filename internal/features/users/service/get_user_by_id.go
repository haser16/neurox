package users_service

import (
	"context"
	"fmt"
	core_domain "neurox/internal/core/domain"
)

func (s *UsersService) GetUserByID(
	ctx context.Context,
	userID int64,
) (core_domain.User, error) {
	userDomain, err := s.usersRepository.GetUserByID(ctx, userID)
	if err != nil {
		return core_domain.User{}, fmt.Errorf(
			"failed to get user by id: %w",
			err,
		)
	}

	if userDomain.Avatar != nil {
		avatarURL, err := s.s3.GetURL(ctx, *userDomain.Avatar)
		if err != nil {
			return core_domain.User{}, fmt.Errorf(
				"failed to get avatar url: %w",
				err,
			)
		}

		userDomain.Avatar = &avatarURL
	}

	return userDomain, nil
}
