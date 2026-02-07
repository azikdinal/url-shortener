package storage

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"shorten/internal/config"
)

func BuildPostgres(cfg *config.Config) (*pgxpool.Config, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)

	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	poolCfg.MaxConns = int32(cfg.Postgres.MaxOpenConn)
	poolCfg.MinConns = int32(cfg.Postgres.MinOpenConn)
	poolCfg.MaxConnLifetime = cfg.Postgres.MaxConnLifetime
	poolCfg.MaxConnIdleTime = cfg.Postgres.MaxConnIdletime

	return poolCfg, nil
}
