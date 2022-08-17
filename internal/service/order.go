package service

import (
	"encoding/json"
	"github.com/onemgvv/wb-l0/internal/domain"
	"github.com/onemgvv/wb-l0/internal/repository"
	"github.com/patrickmn/go-cache"
	"log"
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

func (s *OrderService) GetById(uid string) (domain.OrderJSON, error) {
	var order domain.OrderJSON

	jsonString, ok := s.cache.Get(uid)
	if ok {
		if err := unmarshall(jsonString.(string), &order); err != nil {
			log.Fatalf("[GET ORDER]: %s", err.Error())
			return nil, err
		}

		return order, nil
	}

	ords, err := s.repo.GetOrder(uid)
	if err != nil {
		return nil, err
	}

	if err = unmarshall(ords.Data, &order); err != nil {
		return nil, err
	}

	return order, nil
}

func unmarshall(data string, target *domain.OrderJSON) error {
	return json.Unmarshal([]byte(data), &target)
}
