package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"shorten/internal/app"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGINT,
	)
	defer stop()

	application, err := app.CreateAndRun(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	application.Shutdown(shutdownCtx)
}
