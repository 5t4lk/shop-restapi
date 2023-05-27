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

func (r *ProductMongo) GetAll() ([]types.GetProducts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var products []types.GetProducts

	for cur.Next(ctx) {
		var product types.GetProducts
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductMongo) GetById(userId, productId string) (types.CreateProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return types.CreateProduct{}, err
	}

	filter := bson.M{
		"_id": _id,
	}

	var product types.CreateProduct
	err = r.collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return types.CreateProduct{}, errors.New("list not found")
		}
		return types.CreateProduct{}, err
	}

	return product, nil
}
