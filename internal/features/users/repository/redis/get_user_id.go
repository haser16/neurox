package users_redis_repository

import (
	"context"
	"fmt"
	core_redis_client "neurox/internal/core/repository/redis/client"
)

func (r *UsersRedisRepository) GetUserID(ctx context.Context, token string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.client.GetOperationTimeout())
	defer cancel()

	userID, err := r.client.Get(ctx, token, core_redis_client.VerificationPrefix)
	if err != nil {
		return -1, fmt.Errorf("failed to fetch user id: %w", err)
	}

	return userID, nil
}
