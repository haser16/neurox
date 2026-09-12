package users_postgres_repository

import (
	"context"
	"fmt"
	core_domain "neurox/internal/core/domain"
)

func (r *UsersRepository) GetUserByID(
	ctx context.Context,
	userID int64,
) (core_domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `SELECT
    id, username, email, phone, avatar
	FROM neurox.users WHERE id = $1;`
	row := r.pool.QueryRow(ctx, query, userID)

	var user core_domain.User

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Phone, &user.Avatar)
	if err != nil {
		return user, fmt.Errorf("scan error: %v", err)
	}

	return user, nil
}
