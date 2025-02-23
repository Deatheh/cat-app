package service

import (
	"github.com/Deatheh/cat-app"
	"github.com/Deatheh/cat-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user cat.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type CatList interface {
}

type Cat interface {
}

type Service struct {
	Authorization
	CatList
	Cat
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
