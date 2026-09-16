package requests_postgres_repository

import core_postgres_pool "neurox/internal/core/repository/postgres/pool"

type RequestsRepository struct {
	pool core_postgres_pool.Pool
}

func NewRequestsRepository(pool core_postgres_pool.Pool) *RequestsRepository {
	return &RequestsRepository{pool: pool}
}
