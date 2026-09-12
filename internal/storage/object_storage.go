package storage

import (
	"context"
	"io"
)

type ObjectStorage interface {
	Upload(
		ctx context.Context,
		key string,
		body io.Reader,
		contentType string,
	) error

	Delete(
		ctx context.Context,
		key string,
	) error

	GetURL(key string) string
}
