package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gritsulyak/awesomeProject/internal/application"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := application.NewApp()
	if err := app.Start(ctx); err != nil {
		log.Fatalf("failed to start app: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("shutting down...")
	if err := app.Stop(ctx); err != nil {
		log.Fatalf("failed to stop app: %v", err)
	}
}
