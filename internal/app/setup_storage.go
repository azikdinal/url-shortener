package app

import (
	"context"
	"errors"

	"shorten/internal/config"
	inmemoryStorage "shorten/internal/storage/link/inmemory"
	postgresStorage "shorten/internal/storage/link/postgres"
	useCase "shorten/internal/usecase/link"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Функция для определения типа хранилища и его конфигурации
func setupStorage(ctx context.Context, cfg *config.Config) (
	useCase.LinkStorage,
	*pgxpool.Pool,
	error,
) {

	// Выбирается тип хранилища
	switch cfg.StoreType {

	case "IN_MEMORY":
		return inmemoryStorage.New(), nil, nil

	case "POSTGRES":
		pool, err := connectPostgres(ctx, cfg)
		if err != nil {
			return nil, nil, ErrSetupConnectionPool
		}
		return postgresStorage.New(pool), pool, nil

	default:
		return nil, nil, ErrUnknownStorageType
	}
}

var ErrSetupConnectionPool = errors.New("couldn't setup connection pool to PostgreSQL")
var ErrUnknownStorageType = errors.New("selected unknown storage type")
