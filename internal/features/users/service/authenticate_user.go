package users_service

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (s *UsersService) AuthenticateUser(
	ctx context.Context,
	email string,
	password string,
) (string, error) {
	hashPassword, userID, err := s.usersRepository.GetPasswordByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("failed to get user by `username`: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid password: %w", err)
	}

	token, err := s.tokenService.Generate(userID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return token, nil
}
