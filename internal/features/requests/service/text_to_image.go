package requests_service

import (
	"bytes"
	"context"
	"fmt"
	core_domain "neurox/internal/core/domain"
	integrations_gemini "neurox/internal/integrations/gemini"

	"github.com/google/uuid"
)

func (s *RequestsService) TextToImage(
	ctx context.Context,
	request core_domain.Request,
) (core_domain.Request, error) {
	image, err := s.imageGenerator.Generate(
		ctx,
		request.Prompt,
		integrations_gemini.ModelNanoBananaPro,
	)

	if err != nil {
		return core_domain.Request{}, fmt.Errorf("failed to generate image: %w", err)
	}

	key := fmt.Sprintf(
		"requests/%d/%s",
		request.UserID,
		uuid.NewString(),
	)

	if err := s.s3.Upload(
		ctx,
		key,
		bytes.NewReader(image),
		"image/png",
	); err != nil {
		return core_domain.Request{}, fmt.Errorf("upload avatar to storage: %w", err)
	}

	request.Image = key
	requestDomain, err := s.requestsRepository.TextToImage(ctx, request)
	if err != nil {
		return core_domain.Request{}, fmt.Errorf("save request domain: %w", err)
	}

	requestDomain.Image, err = s.s3.GetURL(ctx, key)
	if err != nil {
		return core_domain.Request{}, fmt.Errorf("get url from storage: %w", err)
	}

	return requestDomain, nil
}
