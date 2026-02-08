package app

import (
	"context"
	"log"
	"net/http"
	"net/url"

	"shorten/internal/config"
	"shorten/internal/router"
	linkHandler "shorten/internal/transport/link/rest"
	linkUseCase "shorten/internal/usecase/link"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
)

type App struct {
	httpServer *http.Server
	grpcServer *grpc.Server
	dbPool     *pgxpool.Pool
}

// Функция для создания и запуска приложения
func CreateAndRun(ctx context.Context) (*App, error) {
	// 1. Загружаются конфиги из .env
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// 2. Настройка хранилища
	linkStorage, dbPool, err := setupStorage(ctx, cfg)
	if err != nil {
		return nil, err
	}

	// 3. Создание Usecase (бизнес-логика)
	pubHost, _ := url.Parse(cfg.PublicHost)
	linkUC := linkUseCase.New(linkStorage, pubHost)

	// 4. Создание handler'а
	linkHandler := linkHandler.New(linkUC, cfg.PublicHost)

	// 5. Конфигурирование HTTP-сервера
	routerEngine := router.NewEngine(linkHandler)

	// 6. Определяется конфигурация GRPC-сервера
	grpcServer, grpcListener, err := buildGRPCServer(linkUC, cfg)

	// 7. Запуск HTTP-сервера
	httpServer := &http.Server{
		Addr:    cfg.HTTP.Address,
		Handler: routerEngine,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("http server error: %v", err)
		}
	}()

	// 8. Запуск GRPC-сервиса
	go func() {
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Println(err)
		}
	}()

	<-ctx.Done()

	return &App{
		httpServer: httpServer,
		grpcServer: grpcServer,
		dbPool:     dbPool,
	}, nil
}
