package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gritsulyak/awesomeProject/internal/application"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	app := application.NewApp()
	if err := app.Start(ctx); err != nil {
		log.Fatalf("failed to start app: %v", err)
	}

	<-done

	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	app.Stop(shutdownCtx)
}
