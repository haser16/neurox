package users_service

import (
	"context"
	"fmt"
	broker_redis "neurox/internal/broker/rabbitmq"
	core_domain "neurox/internal/core/domain"

	"golang.org/x/crypto/bcrypt"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	userDomain core_domain.User,
) (core_domain.User, error) {
	if err := userDomain.Validate(); err != nil {
		return core_domain.User{}, fmt.Errorf("user validation failed: %w", err)
	}

	hashedPassword, err := hashPassword(userDomain.Password)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("hashing password failed: %w", err)
	}

	userDomain.Password = hashedPassword

	userDomain, err = s.usersRepository.CreateUser(ctx, userDomain)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("creating user failed: %w", err)
	}

	if err := s.publisher.Publish(
		ctx,
		"",
		broker_redis.QueueEmailTasks,
		[]byte(userDomain.Email),
	); err != nil {
		return core_domain.User{}, fmt.Errorf("publishing email task failed: %w", err)
	}
	return userDomain, nil
}

func hashPassword(password string) (string, error) {
	passwordBytes := []byte(password)

	hashedBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}

	return string(hashedBytes), nil
}
