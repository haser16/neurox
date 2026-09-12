package broker_redis

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	User     string `envconfig:"USER"`
	Password string `envconfig:"PASSWORD"`
	Host     string `envconfig:"HOST"`
	Port     string `envconfig:"PORT"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("PUBLISHER", &config); err != nil {
		return Config{}, fmt.Errorf("failed to process env: %w", err)
	}
	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err := fmt.Errorf("env publisher read: %w", err)
		panic(err)
	}
	return config
}
