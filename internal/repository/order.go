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

func (r *OrderRepository) Create(id string, data string) (string, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return "", err
	}

	query := fmt.Sprintf("INSERT INTO %s (id, data) VALUES ($1, $2) RETURNING id", ordersTable)

	row := tx.QueryRow(query, id, data)

	if err = row.Err(); err != nil {
		_ = tx.Rollback()
		return "", err
	}

	return id, tx.Commit()
}

func (r *OrderRepository) GetOrder(uid string) (domain.Order, error) {
	var order domain.Order
	query := fmt.Sprint("SELECT * FROM orders WHERE id = $1")
	err := r.db.Get(&order, query, uid)

	return order, err
}
