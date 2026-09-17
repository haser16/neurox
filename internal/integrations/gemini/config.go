package integrations_gemini

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	APIKey  string        `envconfig:"API_KEY"`
	Timeout time.Duration `envconfig:"TIMEOUT"`
}

func NewConfig() (*Config, error) {
	var config Config

	if err := envconfig.Process("GEMINI", &config); err != nil {
		return nil, fmt.Errorf("failed to process env varibles: %w", err)
	}
	return &config, nil
}

func NewConfigMust() *Config {
	config, err := NewConfig()
	if err != nil {
		err := fmt.Errorf("failed to create config gemini: %w", err)
		panic(err)
	}
	return config
}
