package repository

import (
	"fmt"
	"strings"

	"github.com/Deatheh/cat-app"
	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

type CatListPostgres struct {
	db      *sqlx.DB
	minioDb *minio.Client
}

func NewCatListPostgres(db *sqlx.DB, minioDb *minio.Client) *CatListPostgres {
	return &CatListPostgres{
		db:      db,
		minioDb: minioDb,
	}
}

func (r *CatListPostgres) Create(userId int, list cat.CatList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", catListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *CatListPostgres) GetAll(userId int) ([]cat.CatList, error) {
	var lists []cat.CatList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1", catListTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *CatListPostgres) GetById(userId, id int) (cat.CatList, error) {
	var list cat.CatList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND tl.id = $2", catListTable, usersListsTable)
	err := r.db.Get(&list, query, userId, id)
	return list, err
}

func (r *CatListPostgres) Delete(userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND tl.id = $2", catListTable, usersListsTable)
	_, err := r.db.Exec(query, userId, id)
	return err
}

func (r *CatListPostgres) Update(userId, listId int, input cat.UpdeteListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d", catListTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
