package app

import (
	"context"
	"log"
)

func (app *App) Shutdown(ctx context.Context) {
	if err := app.httpServer.Shutdown(ctx); err != nil {
		log.Printf("http shutdown error: %v", err)
	}

	if app.grpcServer != nil {
		app.grpcServer.GracefulStop()
	}

	if app.dbPool != nil {
		app.dbPool.Close()
	}
}
