package users_redis_repository

import (
	"context"
	"fmt"
	core_redis_client "neurox/internal/core/repository/redis/client"
	"time"
)

func (r *UsersRedisRepository) AddUserToken(ctx context.Context, token string, userID int64) error {
	ctx, cancel := context.WithTimeout(ctx, r.client.GetOperationTimeout())
	defer cancel()
	
	if err := r.client.Set(ctx, token, userID, core_redis_client.VerificationPrefix, time.Minute*15); err != nil {
		return fmt.Errorf("failed to add token to redis: %w", err)
	}
	return nil
}
