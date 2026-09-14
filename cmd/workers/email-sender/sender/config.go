package email_sender

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	From     string `env:"FROM"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("SMTP", &config); err != nil {
		return Config{}, fmt.Errorf("process env config: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get jwt config: %w", err)
		panic(err)
	}
	return config
}
