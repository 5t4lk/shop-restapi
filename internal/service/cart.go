package service

import (
	"online_shop/internal/repository"
	"online_shop/internal/types"
)

type CartService struct {
	repo repository.Cart
}

func NewCartService(repo repository.Cart) *CartService {
	return &CartService{repo: repo}
}

func (c *CartService) Add(userId string, cart types.AddToCart) error {
	return c.repo.Add(userId, cart)
}

func (c *CartService) Delete(userId string, cart types.RemoveFromCart) error {
	return c.repo.Delete(userId, cart)
}

func (c *CartService) GetAll(userId string) ([]types.CartProduct, error) {
	return c.repo.GetAll(userId)
}
