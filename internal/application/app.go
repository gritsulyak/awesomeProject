package application

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "net/http/pprof"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"

	"github.com/gritsulyak/awesomeProject/internal/config"
	"github.com/gritsulyak/awesomeProject/internal/model"
	"github.com/gritsulyak/awesomeProject/internal/repository/satellite"
	"github.com/gritsulyak/awesomeProject/internal/service/cache"
	satelliteService "github.com/gritsulyak/awesomeProject/internal/service/satellite"
	v1 "github.com/gritsulyak/awesomeProject/internal/transport/http/v1"
)

type App struct {
	cfg       *config.Config
	db        *sql.DB
	echo      *echo.Echo
	metricsrv *http.Server
}

func NewApp() *App {
	return &App{
		cfg: config.NewConfigFromEnv(),
	}
}

func (app *App) Start(ctx context.Context) error {
	log.Printf("Starting app with config: %+v", app.cfg)
	db, err := openDB(app.cfg.Database.GetDSN())
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	app.db = db

	rdb := redis.NewClient(&redis.Options{
		Addr:     app.cfg.Redis.Addr,
		Password: os.Getenv("REDIS_PASSWORD"), // password ← обязательно!
		DB:       0,
	})

	repo := satellite.NewRepository(app.db)
	satelliteCache := cache.New[*model.Satellite](rdb)
	svc := satelliteService.NewService(repo, satelliteCache)

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiGroup := e.Group("/api/v1/satellite")
	v1.NewController(apiGroup, svc)
	app.echo = e

	go func() {
		if err := e.Start(app.cfg.HttpServerConfig.ListenAddress); err != nil && err != http.ErrServerClosed {
			e.Logger.Errorf("http server error: %v", err)
		}
	}()

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/debug/pprof/", http.DefaultServeMux)
	app.metricsrv = &http.Server{
		Addr:    app.cfg.HttpServerConfig.MetricsAddress,
		Handler: mux,
	}
	go func() {
		if err := app.metricsrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			e.Logger.Errorf("metrics server error: %v", err)
		}
	}()

	return nil
}

func (app *App) Stop(ctx context.Context) error {
	shutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if app.echo != nil {
		_ = app.echo.Shutdown(shutCtx)
	}
	if app.metricsrv != nil {
		_ = app.metricsrv.Shutdown(shutCtx)
	}
	if app.db != nil {
		_ = app.db.Close()
	}
	return nil
}
