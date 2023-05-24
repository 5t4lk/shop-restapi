package shop

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"online_shop/internal/types"
)

type RepoShop struct {
	mongo *mongo.Collection
}

func NewRepoShop(mongo *mongo.Collection) *RepoShop {
	return &RepoShop{mongo: mongo}
}

func (r *RepoShop) Create(ctx context.Context, v types.ProductCreateInput) error {
	return r.Create(ctx, v)
}
