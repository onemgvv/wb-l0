package nats

import (
	"github.com/nats-io/stan.go"
	"github.com/onemgvv/wb-l0/internal/config"
)

func Init(cfg *config.Config) (stan.Conn, error) {
	return stan.Connect(cfg.NATS.ClusterID, cfg.NATS.ClientID, stan.NatsURL("0.0.0.0:4222"))
}
