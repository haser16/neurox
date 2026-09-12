package users_postgres_repository

import (
	"context"
	"fmt"
	core_errors "neurox/internal/core/errors"
)

func (r *UsersRepository) DeleteUser(
	ctx context.Context,
	userID int64,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `DELETE FROM neurox.users WHERE id = $1`
	cmdTag, err := r.pool.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user with id='%d' not found: %w", userID, core_errors.ErrNotFound)
	}
	return nil
}
