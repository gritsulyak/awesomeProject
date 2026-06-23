package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(1)

	pprofSrv := &http.Server{
		Addr:    "127.0.0.1:6060",
		Handler: http.DefaultServeMux, // важно для net/http/pprof [web:114][web:133]
	}

	// Prefer Listen so we can detect "address already in use" early
	ln, err := net.Listen("tcp", pprofSrv.Addr)
	if err != nil {
		log.Fatalf("pprof listen failed on %s: %v", pprofSrv.Addr, err)
	}

	go func() {
		defer wg.Done()
		log.Printf("pprof: http://%s/debug/pprof/", pprofSrv.Addr)

		if err := pprofSrv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("pprof serve error: %v", err)
		}
	}()

	<-done

	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	app.Stop(shutdownCtx)

	if err := pprofSrv.Shutdown(shutdownCtx); err != nil {
		log.Printf("pprof shutdown: %v", err)
	}
	wg.Wait()
}
