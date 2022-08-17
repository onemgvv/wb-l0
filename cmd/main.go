package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
	"github.com/onemgvv/wb-l0/internal/config"
	deliveryHttp "github.com/onemgvv/wb-l0/internal/delivery/http"
	"github.com/onemgvv/wb-l0/internal/logger"
	"github.com/onemgvv/wb-l0/internal/repository"
	"github.com/onemgvv/wb-l0/internal/server"
	"github.com/onemgvv/wb-l0/internal/service"
	"github.com/onemgvv/wb-l0/pkg/database/postgres"
	"github.com/onemgvv/wb-l0/pkg/nats"
	"github.com/patrickmn/go-cache"
	"github.com/tidwall/gjson"
	"log"
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

	c := cache.New(10*time.Minute, 5*time.Minute)

	cfg, err := config.Init(configPath)
	if err != nil {
		logger.ErrorLogger.Errorf("[CONFIG ERROR]: %s\n", err.Error())
	}

	db, err := postgres.Init(cfg)
	if err != nil {
		logger.ErrorLogger.Errorf("[POSTGRES DB ERROR]: %s\n", err.Error())
	}

	if err = postgres.FillCache(db, c); err != nil {
		log.Fatalf("[Fill cache from db ERROR]: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(&service.Deps{
		Repos: repositories,
		Cache: c,
	})
	handler := deliveryHttp.NewHandler(services)
	app := server.NewServer(cfg, handler.InitRoutes())

	go func() {
		if err = app.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("[SERVER START] || [FAILED]: %s", err.Error())
			//logger.ErrorLogger.Fatalf("[SERVER START] || [FAILED]: %s", err.Error())
		}
	}()

	logger.InfoLogger.Infof("Application started on PORT: %s", cfg.HTTP.Port)

	conn, err := nats.Init(cfg)
	if err != nil {
		logger.ErrorLogger.Fatalf("[NATS CONNECTING ERROR]: %s", err.Error())
	}

	sb, err := conn.Subscribe("order", func(m *stan.Msg) {
		fmt.Println("DATA: ", string(m.Data))

		if !gjson.Valid(string(m.Data)) {
			log.Println("JSON not valid")
			return
		}

		id, err := uuid.NewRandom()
		if err != nil {
			log.Fatalf("UUID GEN ERROR: %s", err.Error())
		}

		fmt.Println("ID", id)

		query := fmt.Sprint("INSERT INTO orders (id, data) VALUES ($1, $2) RETURNING id")
		row := db.QueryRow(query, id.String(), string(m.Data))

		if err = row.Scan(&id); err != nil {
			log.Fatalf(err.Error())
		}

		c.Set(id.String(), string(m.Data), 1*time.Hour)
	})

	if err != nil {
		logger.ErrorLogger.Fatalf("[SUB ERROR]: %s", err.Error())
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err = conn.Close(); err != nil {
		logger.ErrorLogger.Fatalf("[NATS CLOSING CONNECTION]: %s", err.Error())
	}

	if err = sb.Unsubscribe(); err != nil {
		logger.ErrorLogger.Fatalf("[NATS UNSUBSCRIBE ERROR]: %s", err.Error())
	}

	if err = sb.Close(); err != nil {
		logger.ErrorLogger.Fatalf("[NATS CLOSING ERROR]: %s", err.Error())
	}

	if err = app.Stop(ctx); err != nil {
		logger.ErrorLogger.Fatalf("[SERVER STOP] || [FAILED]: %s", err.Error())
	}

	if err = postgres.Close(db); err != nil {
		logger.ErrorLogger.Fatalf("[DATABASE CONN CLOSE] || [FAILED]: %s", err.Error())
	}

	logger.InfoLogger.Info("Application stopped")
}
