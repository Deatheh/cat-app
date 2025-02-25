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
	Create(userId int, list cat.CatList) (int, error)
	GetAll(userId int) ([]cat.CatList, error)
	GetById(userId, id int) (cat.CatList, error)
	Delete(userId, id int) error
	Update(userId, id int, input cat.UpdeteListInput) error
}

type Cat interface {
	Create(userId, listId int, list cat.Cat) (int, error)
	GetAll(userId, listId int) ([]cat.Cat, error)
	GetById(userId, listId, itemId int) (cat.Cat, error)
	Delete(userId, itemId int) error
	Update(userId, listId, itemId int, input cat.UpdateCatInput) error
}

type Service struct {
	Authorization
	CatList
	Cat
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		CatList:       NewCatListService(repos.CatList),
		Cat:           NewCatService(repos.Cat, repos.CatList),
	}
}
