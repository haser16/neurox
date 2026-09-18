package users_service

import (
	"context"
	"neurox/internal/auth"
	broker_redis "neurox/internal/broker/rabbitmq"
	core_domain "neurox/internal/core/domain"
	storage_s3 "neurox/internal/storage/s3"
)

type UsersService struct {
	usersRepository      UsersRepository
	usersRedisRepository UsersRedisRepository
	tokenService         auth.TokenService
	s3                   *storage_s3.Storage
	publisher            *broker_redis.Publisher
}

type UsersRepository interface {
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
	GetPasswordByEmail(
		ctx context.Context,
		email string,
	) (string, int64, error)
	DeleteUser(
		ctx context.Context,
		userID int64,
	) error
	UpdateAvatarKey(
		ctx context.Context,
		userID int64,
		key string,
	) error
	EmailVerify(
		ctx context.Context,
		userID int64,
	) error
}

type UsersRedisRepository interface {
	AddUserToken(
		ctx context.Context,
		token string,
		userID int64,
	) error
	GetUserID(
		ctx context.Context,
		token string,
	) (int64, error)
	Delete(
		ctx context.Context,
		token string,
	) error
}

func NewUsersService(
	usersRepository UsersRepository,
	usersRedisRepository UsersRedisRepository,
	tokenService auth.TokenService,
	s3 *storage_s3.Storage,
	publisher *broker_redis.Publisher,
) *UsersService {
	return &UsersService{
		usersRepository:      usersRepository,
		usersRedisRepository: usersRedisRepository,
		tokenService:         tokenService,
		s3:                   s3,
		publisher:            publisher,
	}
}
