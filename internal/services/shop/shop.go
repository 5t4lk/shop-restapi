package shop

import (
	"context"
	"online_shop/internal/repository"
	"online_shop/internal/types"
)

type ServiceShop struct {
	repo repository.Shop
}

func InitServiceShop(repo repository.Shop) *ServiceShop {
	return &ServiceShop{
		repo: repo,
	}
}

func (b *ServiceShop) Create(ctx context.Context, v types.ProductCreateInput) error {
	return b.repo.Create(ctx, v)
}
