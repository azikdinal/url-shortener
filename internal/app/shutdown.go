package app

import (
	"context"
)

func (app *App) Shutdown(ctx context.Context) {
	_ = app.httpServer.Shutdown(ctx)

	if app.grpcServer != nil {
		app.grpcServer.GracefulStop()
	}

	if app.dbPool != nil {
		app.dbPool.Close()
	}
}
