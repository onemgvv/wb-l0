package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/onemgvv/wb-l0/internal/config"
	"github.com/onemgvv/wb-l0/internal/logger"
	"github.com/onemgvv/wb-l0/pkg/database/postgres"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.ErrorLogger.Errorf("[ENV LOAD ERROR]: %s\n", err.Error())
	}

	cfg, err := config.Init("configs")
	if err != nil {
		logger.ErrorLogger.Errorf("[CONFIG ERROR]: %s\n", err.Error())
	}

	db, err := postgres.Init(cfg)
	if err != nil {
		logger.ErrorLogger.Errorf("[POSTGRES DB ERROR]: %s\n", err.Error())
	}

	uid, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("UUID GEN ERROR: %s", err.Error())
	}

	data := `{"orderID": "542", "sum": "1.300", "curr": "RUB"}`

	var id string

	row := db.QueryRow("INSERT INTO orders (id, data) VALUES ($1, $2) RETURNING id", uid, data)
	if err := row.Scan(&id); err != nil {
		fmt.Println("SCAN ERROR", err.Error())
	}

	fmt.Println(id)
}
