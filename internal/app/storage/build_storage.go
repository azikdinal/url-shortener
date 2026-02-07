package storage

import (
	"errors"

	"shorten/internal/config"
	inmemoryStorage "shorten/internal/storage/link/inmemory"
	postgresStorage "shorten/internal/storage/link/postgres"
	useCase "shorten/internal/usecase/link"

	"github.com/jackc/pgx/v4/pgxpool"
)

func BuildStorage(cfg *config.Config) (useCase.LinkStorage, *pgxpool.Config, *pgxpool.Pool, error) {
	// Выбирается тип хранилища
	switch cfg.StoreType {

	case "IN_MEMORY":
		return inmemoryStorage.New(), nil, nil, nil

	case "POSTGRES":
		poolCfg, err := BuildPostgres(cfg)
		if err != nil {
			return nil, nil, nil, ErrSetupConnectionPool
		}
		emptyPool := &pgxpool.Pool{}
		return postgresStorage.New(emptyPool), poolCfg, emptyPool, nil

	default:
		return nil, nil, nil, ErrUnknownStorageType
	}
}

var ErrSetupConnectionPool = errors.New("couldn't setup connection pool to PostgreSQL")
var ErrUnknownStorageType = errors.New("selected unknown storage type")
