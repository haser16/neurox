package requests_postgres_repository

import (
	"context"
	"fmt"
	core_domain "neurox/internal/core/domain"
)

func (r *RequestsRepository) TextToImage(
	ctx context.Context,
	requests core_domain.Request,
) (core_domain.Request, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `INSERT INTO
    neurox.requests (prompt, image, user_id)
	VALUES 
    ($1, $2, $3)
    RETURNING id, prompt, image;`

	row := r.pool.QueryRow(
		ctx,
		query,
		requests.Prompt,
		requests.Image,
		requests.UserID,
	)
	var request core_domain.Request

	err := row.Scan(&request.ID, &request.Prompt, &request.Image)
	if err != nil {
		return request, fmt.Errorf("scan error: %w", err)
	}
	return request, nil
}
