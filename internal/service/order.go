package service

import (
	"online_shop/internal/repository"
	"online_shop/internal/types"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (q *OrderService) Place(userID string, cart types.ShoppingCart) (string, error) {
	return q.repo.Place(userID, cart)
}

func (q *OrderService) Delete(userId, orderId string) error {
	return q.repo.Delete(userId, orderId)
}
