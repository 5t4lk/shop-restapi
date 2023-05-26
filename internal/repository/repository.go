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

type Repository struct {
	Authorization
}

func NewRepository(client *mongo.Client, dbName string) *Repository {
	db := client.Database(dbName)

	return &Repository{
		Authorization: repository.NewAuthMongo(db.Collection("users")),
	}
}
