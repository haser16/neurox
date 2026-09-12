package users_postgres_repository

import (
	"context"
	"fmt"
	core_domain "neurox/internal/core/domain"
)

func (r *UsersRepository) GetUserByUsername(
	ctx context.Context,
	username string,
) (core_domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `SELECT id, username, email, phone FROM neurox.users WHERE username = $1`

	row := r.pool.QueryRow(ctx, query, username)

	var user core_domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Phone)

	if err != nil {
		return user, fmt.Errorf("scan error: %w", err)
	}

	return user, nil
}
