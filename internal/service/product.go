package service

import (
	"online_shop/internal/repository"
	"online_shop/internal/types"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(userId string, product types.CreateProduct) (string, error) {
	return s.repo.Create(userId, product)
}
