package core_redis_client

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	client    *redis.Client
	opTimeout time.Duration
}

func NewClient(config *Config) *Client {
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	client := redis.NewClient(&redis.Options{Addr: addr})
	return &Client{client: client, opTimeout: config.Timeout}
}
