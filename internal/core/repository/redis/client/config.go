package core_redis_client

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host    string        `envconfig:"HOST" required:"true"`
	Port    int           `envconfig:"PORT" required:"true"`
	Timeout time.Duration `envconfig:"TIMEOUT" required:"true"`
}

func NewConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("REDIS", &config); err != nil {
		return nil, fmt.Errorf("failed to process env vars: %w", err)
	}
	return &config, nil
}

func NewConfigMust() *Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("load env redis config: %w", err)
		panic(err)
	}
	return config
}
