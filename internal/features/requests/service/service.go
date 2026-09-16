package requests_service

import (
	"context"
	core_domain "neurox/internal/core/domain"
	integrations_gemini "neurox/internal/integrations/gemini"
	storage_s3 "neurox/internal/storage/s3"
)

type RequestsService struct {
	requestsRepository RequestsRepository
	imageGenerator     integrations_gemini.Generator
	s3                 *storage_s3.Storage
}

type RequestsRepository interface {
	TextToImage(
		ctx context.Context,
		requests core_domain.Request,
	) (core_domain.Request, error)
}

func NewRequestsService(
	requestsRepository RequestsRepository,
	imageGenerator integrations_gemini.Generator,
	s3 *storage_s3.Storage,
) *RequestsService {
	return &RequestsService{
		requestsRepository: requestsRepository,
		imageGenerator:     imageGenerator,
		s3:                 s3,
	}
}
