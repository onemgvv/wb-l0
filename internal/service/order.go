package service

import (
	"github.com/onemgvv/wb-l0/internal/domain"
	"github.com/onemgvv/wb-l0/internal/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetById(uid string) (domain.Order, error) {
	return s.repo.GetOrder(uid)
}
