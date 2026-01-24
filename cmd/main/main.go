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
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		if err := app.Run(); err != nil {
			log.Println(err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	app.Shutdown(shutdownCtx)
}
