package main

import (
	"context"
	"errors"
	"github.com/joho/godotenv"
	"github.com/onemgvv/wb-l0/internal/config"
	deliveryHttp "github.com/onemgvv/wb-l0/internal/delivery/http"
	"github.com/onemgvv/wb-l0/internal/logger"
	"github.com/onemgvv/wb-l0/internal/repository"
	"github.com/onemgvv/wb-l0/internal/server"
	"github.com/onemgvv/wb-l0/internal/service"
	"github.com/onemgvv/wb-l0/pkg/database/postgres"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const configPath = "configs"

func main() {
	if err := godotenv.Load(); err != nil {
		logger.ErrorLogger.Errorf("[ENV LOAD ERROR]: %s\n", err.Error())
	}

	cfg, err := config.Init(configPath)
	if err != nil {
		logger.ErrorLogger.Errorf("[CONFIG ERROR]: %s\n", err.Error())
	}

	db, err := postgres.Init(cfg)
	if err != nil {
		logger.ErrorLogger.Errorf("[POSTGRES DB ERROR]: %s\n", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(&service.Deps{
		Repos: repositories,
	})
	handler := deliveryHttp.NewHandler(services)
	app := server.NewServer(cfg, handler.InitRoutes())

	go func() {
		if err = app.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorLogger.Fatalf("[SERVER START] || [FAILED]: %s", err.Error())
		}
	}()

	logger.InfoLogger.Infof("Application started on PORT: %s", cfg.HTTP.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err = app.Stop(ctx); err != nil {
		logger.ErrorLogger.Fatalf("[SERVER STOP] || [FAILED]: %s", err.Error())
	}

	if err = postgres.Close(db); err != nil {
		logger.ErrorLogger.Fatalf("[DATABASE CONN CLOSE] || [FAILED]: %s", err.Error())
	}

	logger.InfoLogger.Info("Application stopped")
}
