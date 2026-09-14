package core_redis_client

import (
	"context"
	"strconv"
	"time"
)

func (s *Client) Set(
	ctx context.Context,
	token string,
	userID int64,
	prefix string,
	ttl time.Duration,
) error {
	return s.client.Set(
		ctx,
		prefix+token,
		userID,
		ttl,
	).Err()
}

func (s *Client) Get(
	ctx context.Context,
	token string,
	prefix string,
) (int64, error) {
	value, err := s.client.Get(
		ctx,
		prefix+token,
	).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(value, 10, 64)
}

func (s *Client) Delete(
	ctx context.Context,
	token string,
	prefix string,
) error {
	return s.client.Del(
		ctx,
		prefix+token,
	).Err()
}

func (s *Client) GetOperationTimeout() time.Duration {
	return s.opTimeout
}
