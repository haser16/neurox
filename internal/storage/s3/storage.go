package storage_s3

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	bucket        string
	publicURL     string
}

func NewStorage(
	client *s3.Client,
	config Config,
) *Storage {
	return &Storage{
		client:        client,
		presignClient: s3.NewPresignClient(client),
		bucket:        config.Bucket,
		publicURL:     strings.TrimRight(config.PublicURL, "/"),
	}
}

func (s *Storage) Upload(
	ctx context.Context,
	key string,
	body io.Reader,
	contentType string,
) error {
	_, err := s.client.PutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket:      aws.String(s.bucket),
			Key:         aws.String(key),
			Body:        body,
			ContentType: aws.String(contentType),
		},
	)

	return err
}

func (s *Storage) Delete(
	ctx context.Context,
	key string,
) error {
	_, err := s.client.DeleteObject(
		ctx,
		&s3.DeleteObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(key),
		},
	)

	return err
}

func (s *Storage) GetURL(
	ctx context.Context,
	key string,
) (string, error) {
	if key == "" {
		return "", nil
	}

	result, err := s.presignClient.PresignGetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(key),
		},
		s3.WithPresignExpires(time.Hour),
	)
	if err != nil {
		return "", fmt.Errorf("failed to create presigned url: %w", err)
	}

	return result.URL, nil
}
