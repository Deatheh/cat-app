package repository

import (
	"github.com/Deatheh/cat-app"
	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
)

type Authorization interface {
	CreateUser(user cat.User) (int, error)
	GetUser(username, password string) (cat.User, error)
}

type CatList interface {
	Create(userId int, list cat.CatList) (int, error)
	GetAll(userId int) ([]cat.CatList, error)
	GetById(userId, id int) (cat.CatList, error)
	Delete(userId, id int) error
	Update(userId, id int, input cat.UpdeteListInput) error
}

type Cat interface {
	Create(listId int, item cat.Cat) (int, error)
	GetAll(userId, listId int) ([]cat.Cat, error)
	GetById(userId, listId, itemId int) (cat.Cat, error)
	Delete(userId, itemId int) error
	Update(userId, listId, itemId int, input cat.UpdateCatInput) error
}

type Repository struct {
	Authorization
	CatList
	Cat
}

func NewRepository(db *sqlx.DB, minioDb *minio.Client) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		CatList:       NewCatListPostgres(db, minioDb),
		Cat:           NewCatPostgres(db, minioDb),
	}
}
