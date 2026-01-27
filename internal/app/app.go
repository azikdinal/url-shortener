package app

import (
	"context"
	"database/sql"
	"log"
	"net"

	"shorten/internal/config"
	"shorten/internal/router"
	"shorten/internal/usecase"

	"google.golang.org/grpc"
)

type App struct {
	httpRouter *router.Router
	grpcServer *grpc.Server
	grpcLis    net.Listener
	db         *sql.DB
	addr       string
}

func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	repo, db, err := BuildRepo(cfg)
	if err != nil {
		return nil, err
	}

	uc := usecase.New(repo)
	httpRouter := BuildHTTP(uc, cfg)
	grpcServer, grpcLis, err := BuildGRPC(uc, cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		httpRouter: httpRouter,
		grpcServer: grpcServer,
		grpcLis:    grpcLis,
		db:         db,
		addr:       cfg.HTTP.Address,
	}, nil
}

func (app *App) Run() error {
	go func() {
		if err := app.grpcServer.Serve(app.grpcLis); err != nil {
			log.Println(err)
		}
	}()
	return app.httpRouter.Run(app.addr)
}

func (app *App) Shutdown(ctx context.Context) {
	_ = app.httpRouter.Shutdown(ctx)

	if app.grpcServer != nil {
		app.grpcServer.GracefulStop()
	}

	if app.db != nil {
		_ = app.db.Close()
	}
}
