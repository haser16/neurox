package users_postgres_repository

import (
	"context"
	"fmt"
)

func (r *UsersRepository) EmailVerify(
	ctx context.Context,
	userID int64,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `
        UPDATE neurox.users
        SET email_verified = TRUE
        WHERE id = $1;
    `

	result, err := r.pool.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("email verify: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user with id %d not found", userID)
	}

	return nil
}
