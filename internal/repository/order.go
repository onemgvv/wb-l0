package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/onemgvv/wb-l0/internal/domain"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetOrder(uid string) (domain.Order, error) {
	var order domain.Order
	query := fmt.Sprint("SELECT * FROM orders WHERE id = $1")
	err := r.db.Get(&order, query, uid)

	return order, err
}
