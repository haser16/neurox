package storage_s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewClient(
	ctx context.Context,
	configs3 Config,
) (*s3.Client, error) {
	awsConfig, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(configs3.Region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				configs3.AccessKey,
				configs3.SecretKey,
				"",
			),
		),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(
		awsConfig,
		func(options *s3.Options) {
			options.BaseEndpoint = aws.String(configs3.Endpoint)
			options.UsePathStyle = true
		},
	)

	return client, nil
}
