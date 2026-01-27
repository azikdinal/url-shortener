package app

import (
	"database/sql"
	"errors"
	"fmt"

	"shorten/internal/config"
	"shorten/internal/repo"
)

func BuildRepo(cfg *config.Config) (repo.Repo, *sql.DB, error) {
	// Выбирается тип хранилища
	switch cfg.StoreType {

	case "IN_MEMORY":
		return repo.NewMemoryRepo(), nil, nil

	case "POSTGRES":
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.User,
			cfg.Postgres.Password,
			cfg.Postgres.DBName,
			cfg.Postgres.SSLMode,
		)
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			return nil, nil, err
		}

		if err = db.Ping(); err != nil {
			return nil, nil, err
		}

		return repo.NewPostgresRepo(db), db, nil

	default:
		return nil, nil, errors.New("selected unknown storage type")
	}
}
