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

func (s *ProductService) GetAll() ([]types.GetProducts, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetById(userId, productId string) (types.CreateProduct, error) {
	return s.repo.GetById(userId, productId)
}

func (s *ProductService) Delete(userId, productId string) error {
	return s.repo.Delete(userId, productId)
}

func (s *ProductService) Update(userId, productId string, input types.UpdateProduct) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, productId, input)
}
