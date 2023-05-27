package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"online_shop/internal/types"
	"time"
)

type ProductMongo struct {
	collection *mongo.Collection
}

func NewProductMongo(collection *mongo.Collection) *ProductMongo {
	return &ProductMongo{collection: collection}
}

func (r *ProductMongo) Create(userId string, product types.CreateProduct) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := map[string]interface{}{
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"stock":       product.Stock,
		"user_id":     userId,
	}

	result, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return "", err
	}

	productId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("failed to get inserted product id")
	}

	return productId.Hex(), nil
}
