package service

import (
	"online_shop/internal/repository"
	"online_shop/internal/types"
)

type Authorization interface {
	CreateUser(user types.User) (string, error)
	GenerateToken(username, password string) (string, error)
	Validate(token string) (string, error)
}

type Product interface {
	Create(userId string, product types.CreateProduct) (string, error)
	GetAll() ([]types.GetProducts, error)
	GetById(userId, productId string) (types.CreateProduct, error)
}

type Service struct {
	Authorization
	Product
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Product:       NewProductService(repos.Product),
	}
}
