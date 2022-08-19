package service

import (
	"github.com/onemgvv/wb-l0/internal/domain"
	"github.com/onemgvv/wb-l0/internal/repository"
	"github.com/patrickmn/go-cache"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Orders interface {
	GetById(uid string) (domain.OrderJSON, error)
}

type Service struct {
	Orders
}

type Deps struct {
	Repos *repository.Repository
	Cache *cache.Cache
}

func NewService(deps *Deps) *Service {
	orderService := NewOrderService(deps)
	return &Service{
		Orders: orderService,
	}
}
