package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"shorten/internal/config"
	"shorten/internal/handler"
	"shorten/internal/repo"
	"shorten/internal/router"
	"shorten/internal/usecase"
)

type App struct {
	router *router.Router
	db     *sql.DB
	addr   string
}

func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	var db *sql.DB
	var shortenerRepo repo.Repo

	// Выбирается тип хранилища
	switch cfg.StoreType {

	case "IN_MEMORY":
		shortenerRepo = repo.NewMemoryRepo()

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
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			return nil, err
		}

		if err := db.Ping(); err != nil {
			return nil, err
		}

		shortenerRepo = repo.NewPostgresRepo(db)

	default:
		return nil, errors.New("selected unknown storage type")
	}

	uc := usecase.New(shortenerRepo)
	h := handler.New(uc, cfg.HTTP.PublicHost)
	r := router.New(h)

	return &App{
		router: r,
		db:     db,
		addr:   cfg.HTTP.Addr,
	}, nil
}

func (app *App) Run() error {
	return app.router.Run(app.addr)
}

func (app *App) Shutdown(ctx context.Context) {
	_ = app.router.Shutdown(ctx)

	if app.db != nil {
		_ = app.db.Close()
	}
}
