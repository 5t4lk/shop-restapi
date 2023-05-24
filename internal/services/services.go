package services

import (
	"context"
	"online_shop/internal/repository"
	"online_shop/internal/services/shop"
	"online_shop/internal/types"
)

type Shop interface {
	Create(ctx context.Context, v types.ProductCreateInput) error
}

type Service struct {
	Shop
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Shop: shop.InitServiceShop(repo.Shop),
	}
}
