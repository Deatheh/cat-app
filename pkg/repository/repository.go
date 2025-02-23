package repository

import (
	"github.com/Deatheh/cat-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user cat.User) (int, error)
	GetUser(username, password string) (cat.User, error)
}

type CatList interface {
}

type Cat interface {
}

type Repository struct {
	Authorization
	CatList
	Cat
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
