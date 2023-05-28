package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	repository "online_shop/internal/database/mongo"
	"online_shop/internal/types"
)

type Authorization interface {
	CreateUser(user types.User) (string, error)
	GetUser(username, password string) (types.User, error)
}

type Product interface {
	Create(userId string, product types.CreateProduct) (string, error)
	GetAll() ([]types.GetProducts, error)
	GetById(userId, productId string) (types.CreateProduct, error)
	Delete(userId, productId string) error
	Update(userId, productId string, input types.UpdateProduct) error
}

type Cart interface {
	Add(userId string, cart types.AddToCart) error
	Delete(userId string, cart types.RemoveFromCart) error
	GetAll(userId string) ([]types.CartProduct, error)
}

type Order interface {
	Place(userID string, cart types.ShoppingCart) (string, error)
	Delete(userId, orderId string) error
}

type Repository struct {
	Authorization
	Product
	Cart
	Order
}

func NewRepository(client *mongo.Client, dbName string) *Repository {
	db := client.Database(dbName)

	return &Repository{
		Authorization: repository.NewAuthMongo(db.Collection("users")),
		Product:       repository.NewProductMongo(db.Collection("products")),
		Cart:          repository.NewCartMongo(db.Collection("shopping_carts")),
		Order:         repository.NewOrderMongo(db.Collection("orders")),
	}
}
