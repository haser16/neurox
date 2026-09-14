package users_redis_repository

import core_redis "neurox/internal/core/repository/redis"

type UsersRedisRepository struct {
	client core_redis.Client
}

func NewUsersRedisRepository(client core_redis.Client) *UsersRedisRepository {
	return &UsersRedisRepository{
		client: client,
	}
}
