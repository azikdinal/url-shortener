package app

import (
	"context"
	"fmt"
	"shorten/internal/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func connectPostgres(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	//////////////////
	// Конфигурация
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

	//////////////////
	// Подключение
	pool, err := pgxpool.ConnectConfig(ctx, poolCfg)
	if err != nil {
		return nil, err
	}

	// Проверяем соединение пингом
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
