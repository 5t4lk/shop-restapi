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

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
