package requests_postgres_repository

import (
	"context"
	"fmt"
	core_domain "neurox/internal/core/domain"
)

func (r *RequestsRepository) GetRequests(
	ctx context.Context,
	limit *int,
	offset *int,
	userID int64,
) ([]core_domain.Request, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `SELECT id, image, prompt
FROM neurox.requests
WHERE user_id = $3
ORDER BY id
LIMIT $1 OFFSET $2;`

	rows, err := r.pool.Query(ctx, query, limit, offset, userID)
	if err != nil {
		return nil, fmt.Errorf("select requests: %w", err)
	}

	defer rows.Close()

	var requests []core_domain.Request
	for rows.Next() {
		var request core_domain.Request
		err := rows.Scan(&request.ID, &request.Image, &request.Prompt)
		if err != nil {
			return nil, fmt.Errorf("scan request: %w", err)
		}
		requests = append(requests, request)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows: %w", err)
	}
	return requests, nil
}
