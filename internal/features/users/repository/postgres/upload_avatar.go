package users_postgres_repository

import (
	"context"
	"fmt"
)

func (r *UsersRepository) UpdateAvatarKey(
	ctx context.Context,
	userID int64,
	key string,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `UPDATE neurox.users SET avatar = $1 WHERE id = $2`

	result, err := r.pool.Exec(ctx, query, key, userID)
	if err != nil {
		return fmt.Errorf("update user avatar: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user with id %d not found", userID)
	}

	return nil
}
