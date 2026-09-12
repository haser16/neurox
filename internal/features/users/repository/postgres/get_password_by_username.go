package users_postgres_repository

import (
	"context"
	"fmt"
)

func (r *UsersRepository) GetPasswordByUsername(
	ctx context.Context,
	username string,
) (string, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `
        SELECT id, password
        FROM neurox.users
        WHERE username = $1
    `

	row := r.pool.QueryRow(ctx, query, username)

	var (
		password string
		userID   int64
	)

	err := row.Scan(&userID, &password)
	if err != nil {
		return "", 0, fmt.Errorf("failed to scan user: %w", err)
	}

	return password, userID, nil
}
