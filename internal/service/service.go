package service

import (
	"github.com/onemgvv/wb-l0/internal/domain"
	"github.com/onemgvv/wb-l0/internal/repository"
)

type Orders interface {
	GetById(uid string) (domain.Order, error)
}

type Service struct {
	Orders
}

type Deps struct {
	Repos *repository.Repository
}

func NewService(deps *Deps) *Service {
	orderService := NewOrderService(deps.Repos.Order)
	return &Service{
		Orders: orderService,
	}
}
