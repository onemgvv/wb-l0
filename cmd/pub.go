package main

import (
	"github.com/joho/godotenv"
	"github.com/onemgvv/wb-l0/internal/config"
	"github.com/onemgvv/wb-l0/internal/logger"
	"github.com/onemgvv/wb-l0/pkg/nats"
	"io/ioutil"
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

	cfg.NATS.ClientID = "publisher"

	data, err := ioutil.ReadFile("model.json")
	if err != nil {
		logger.ErrorLogger.Fatalf("[model.json READ FAILED]: %s", err.Error())
	}

	conn, err := nats.Init(cfg)
	if err != nil {
		log.Fatalf("[NATS CONNECTING ERROR]: %s", err.Error())
	}

	err = conn.Publish("order", data)
	if err != nil {
		log.Fatalf("Pub error %s", err.Error())
	}
}
