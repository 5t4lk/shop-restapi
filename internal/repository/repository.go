package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"online_shop/internal/repository/shop"
	"online_shop/internal/types"
)

type Shop interface {
	Create(ctx context.Context, v types.ProductCreateInput) error
}

type Repository struct {
	Shop
}

func New(client *mongo.Client, dbName string) *Repository {
	db := client.Database(dbName)

	return &Repository{
		Shop: shop.NewRepoShop(db.Collection("products")),
	}
}
