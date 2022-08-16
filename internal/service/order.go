package service

import (
	"github.com/onemgvv/wb-l0/internal/domain"
	"github.com/onemgvv/wb-l0/internal/repository"
	"github.com/patrickmn/go-cache"
)

type OrderService struct {
	repo  repository.Order
	cache *cache.Cache
}

func NewOrderService(deps *Deps) *OrderService {
	return &OrderService{
		repo:  deps.Repos.Order,
		cache: deps.Cache,
	}
}

func (s *OrderService) GetById(uid string) (domain.Order, error) {
	order, ok := s.cache.Get(uid)
	if ok {
		return order.(domain.Order), nil
	}
	return s.repo.GetOrder(uid)
}
