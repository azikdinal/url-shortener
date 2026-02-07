package app

import (
	"net"

	"shorten/internal/app/storage"
	"shorten/internal/app/transport"
	"shorten/internal/config"

	domain "shorten/internal/domain/link"
	"shorten/internal/router"
	linkUseCase "shorten/internal/usecase/link"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
)

type App struct {
	httpServer   *router.Router
	grpcServer   *grpc.Server
	grpcListener net.Listener
	dbPoolCfg    *pgxpool.Config // используется в Run()
	dbPool       *pgxpool.Pool
	httpAddr     string
}

// New() реализует загрузку конфигов компонентов приложения
func New() (*App, error) {
	// Загружаются конфиги из .env
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// Определяется тип хранилища и его конфигурация
	// dbCfg == nil, если выбран inmemory
	linkStorage, dbPoolCfg, emptyDBPool, err := storage.BuildStorage(cfg)
	if err != nil {
		return nil, err
	}

	scGen := domain.NewShortCodeGenerator(cfg.ShortCodeSeed)
	linkUC := linkUseCase.New(linkStorage, scGen)
	// Определяется конфигурация HTTP-сервера
	httpServer := transport.BuildHTTPServer(linkUC, cfg)

	// Определяется конфигурация GRPC-сервера
	grpcServer, grpcListener, err := transport.BuildGRPCServer(linkUC, cfg)

	return &App{
		httpServer:   httpServer,
		grpcServer:   grpcServer,
		grpcListener: grpcListener,
		dbPoolCfg:    dbPoolCfg,
		dbPool:       emptyDBPool,
		httpAddr:     cfg.HTTP.Address,
	}, nil
}
