package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/onemgvv/wb-l0/internal/config"
	"github.com/onemgvv/wb-l0/internal/domain"
	"github.com/patrickmn/go-cache"
	"time"
)

func Init(cfg *config.Config) (*sqlx.DB, error) {
	postgres := cfg.DB.Postgres
	var dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s  password=%s sslmode=%s",
		postgres.Host, postgres.Port, postgres.User, postgres.Name, postgres.Password, postgres.SSLMode)

	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func FillCache(db *sqlx.DB, c *cache.Cache) error {
	var orders []domain.Order

	query := fmt.Sprintf("SELECT * FROM orders")
	if err := db.Select(&orders, query); err != nil {
		return err
	}

	for _, v := range orders {
		c.Set(v.UID, v.Data, 1*time.Hour)
	}

	return nil
}

func Close(db *sqlx.DB) error {
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}
