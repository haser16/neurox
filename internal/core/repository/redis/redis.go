package core_redis

import (
	"context"
	"time"
)

type Client interface {
	Set(
		ctx context.Context,
		token string,
		userID int64,
		prefix string,
		ttl time.Duration,
	) error

	Get(
		ctx context.Context,
		token string,
		prefix string,
	) (int64, error)

	Delete(
		ctx context.Context,
		token string,
		prefix string,
	) error
	GetOperationTimeout() time.Duration
}
