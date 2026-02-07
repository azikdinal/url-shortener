package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresStorage struct {
	dbPool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{
		dbPool: pool,
	}
}
