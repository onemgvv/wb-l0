package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/onemgvv/wb-l0/internal/domain"
)

const (
	ordersTable = "orders"
)

type Order interface {
	Create(id string, data string) (string, error)
	GetOrder(uid string) (domain.Order, error)
}

type Repository struct {
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order: NewOrderRepository(db),
	}
}
