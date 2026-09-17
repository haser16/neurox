package users_redis_repository

import (
	"context"
	"fmt"
	core_redis_client "neurox/internal/core/repository/redis/client"
)

func (r *UsersRedisRepository) Delete(
	ctx context.Context,
	token string,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.client.GetOperationTimeout())
	defer cancel()

	if err := r.client.Delete(ctx, token, core_redis_client.VerificationPrefix); err != nil {
		return fmt.Errorf("failed to delete token from redis: %w", err)
	}
	return nil
}
