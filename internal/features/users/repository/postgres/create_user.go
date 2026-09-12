package users_postgres_repository

import (
	"context"
	"fmt"
	core_domain "neurox/internal/core/domain"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	userDomain core_domain.User,
) (core_domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `INSERT INTO neurox.users 
    (username, email, phone, password) 
	VALUES 
    ($1, $2, $3, $4)
     RETURNING id, username, email, phone;`
	row := r.pool.QueryRow(
		ctx,
		query,
		userDomain.Username,
		userDomain.Email,
		userDomain.Phone,
		userDomain.Password,
	)

	var user core_domain.User

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Phone)
	if err != nil {
		return user, fmt.Errorf("scan error: %w", err)
	}
	return user, nil
}
