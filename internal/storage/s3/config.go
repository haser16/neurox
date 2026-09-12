package storage_s3

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Endpoint  string `envconfig:"ENDPOINT"`
	Region    string `envconfig:"REGION"`
	PublicURL string `envconfig:"PUBLIC_URL"`
	AccessKey string `envconfig:"ACCESS_KEY"`
	SecretKey string `envconfig:"SECRET_KEY"`
	Bucket    string `envconfig:"BUCKET"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("S3", &config); err != nil {
		return Config{}, fmt.Errorf("failed to process S3 env: %w ", err)
	}
	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err := fmt.Errorf("get jwt config: %w ", err)
		panic(err)
	}
	return config
}
