package requests_service

import (
	"context"
	"fmt"
	core_domain "neurox/internal/core/domain"
)

func (s *RequestsService) GetRequests(
	ctx context.Context,
	limit *int,
	offset *int,
	userID int64,
) ([]core_domain.Request, error) {
	requests, err := s.requestsRepository.GetRequests(ctx, limit, offset, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get requests from repository: %w", err)
	}
	for i := range requests {
		requests[i].Image, err = s.s3.GetURL(ctx, requests[i].Image)
		if err != nil {
			return nil, fmt.Errorf("failed to get s3 url: %w", err)
		}
	}
	return requests, nil
}
