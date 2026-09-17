package users_service

import (
	"context"
	"fmt"
)

func (s *UsersService) EmailVerify(
	ctx context.Context,
	token string,
) error {
	userID, err := s.usersRedisRepository.GetUserID(ctx, token)
	if err != nil {
		return fmt.Errorf("failed to get userID from redis: %w", err)
	}

	if err := s.usersRepository.EmailVerify(ctx, userID); err != nil {
		return fmt.Errorf("failed to verify email: %w", err)
	}

	if err := s.usersRedisRepository.Delete(ctx, token); err != nil {
		return fmt.Errorf("failed to delete verification token: %w", err)
	}

	return nil
}
