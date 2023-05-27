package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"online_shop/internal/types"
	"time"
)

type CartMongo struct {
	collection *mongo.Collection
}

func NewCartMongo(collection *mongo.Collection) *CartMongo {
	return &CartMongo{collection: collection}
}

func (n *CartMongo) Add(userId string, cart types.AddToCart) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := types.CartProduct{
		ProductID: cart.ProductID,
		Quantity:  cart.Quantity,
		UserId:    userId,
	}

	_, err := n.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	return nil
}

func (n *CartMongo) Delete(userId string, cart types.RemoveFromCart) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := map[string]interface{}{
		"product_id": cart.ProductID,
	}

	_, err := n.collection.DeleteOne(ctx, doc)
	if err != nil {
		return err
	}

	return nil
}
