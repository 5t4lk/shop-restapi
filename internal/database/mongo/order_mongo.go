package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"online_shop/internal/types"
	"time"
)

type OrderMongo struct {
	collection *mongo.Collection
}

func NewOrderMongo(collection *mongo.Collection) *OrderMongo {
	return &OrderMongo{collection: collection}
}

func (o *OrderMongo) Place(userID string, cart types.ShoppingCart) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order := types.Order{
		UserID:     userID,
		ProductIDs: getProductIDs(cart.Items),
	}

	result, err := o.collection.InsertOne(ctx, order)
	if err != nil {
		return "", err
	}

	orderId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("failed to get inserted product id")
	}

	return orderId.Hex(), nil
}

func getProductIDs(items []types.CartProduct) []string {
	productIDs := make([]string, len(items))
	for i, item := range items {
		productIDs[i] = item.ProductID
	}
	return productIDs
}

func (o *OrderMongo) Delete(userId, orderId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": _id,
	}

	_, err = o.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New("error occurred while deleting list")
	}

	return nil
}
