package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func (app *App) Run(ctx context.Context) error {
	// Postgres
	if app.dbPoolCfg != nil {
		if err := app.connectPostgres(ctx); err != nil {
			return err
		}
	}

	// HTTP REST API
	if err := app.httpServer.Run(app.httpAddr); err != nil {
		return err
	}

	// GRPC
	go app.runGRPC()

	return nil
}

func (app *App) runGRPC() {
	if err := app.grpcServer.Serve(app.grpcListener); err != nil {
		log.Println(err)
	}
}

func (app *App) connectPostgres(ctx context.Context) error {
	// Создаем пулл соединений
	pool, err := pgxpool.ConnectConfig(ctx, app.dbPoolCfg)
	if err != nil {
		return err
	}

	// Проверяем соединение пингом
	if err := pool.Ping(ctx); err != nil {
		return err
	}

	// Заменяем ссылку на пустой пул postgresStorage.pool
	app.dbPool = pool

	return nil
}
