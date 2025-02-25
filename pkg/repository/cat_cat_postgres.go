package repository

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Deatheh/cat-app"
	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
)

type CatPostgres struct {
	db      *sqlx.DB
	minioDb *minio.Client
}

func NewCatPostgres(db *sqlx.DB, minioDb *minio.Client) *CatPostgres {
	return &CatPostgres{
		db:      db,
		minioDb: minioDb,
	}
}

func (r *CatPostgres) Create(listId int, item cat.Cat) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (age, name, filename, description) values ($1, $2, $3, $4) RETURNING id", catsTable)

	row := tx.QueryRow(createItemQuery, item.Age, item.Name, item.FileName, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	CreateBucket(r.minioDb, strconv.Itoa(itemId)+strconv.Itoa(listId)+"bucket")
	err = RFPutObject(r.minioDb, strconv.Itoa(itemId)+strconv.Itoa(listId)+"bucket", item.FileName+".png", item.FileName+".png", "image/png")
	if err != nil {
		return 0, err
	}
	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, cat_id) values ($1, $2)", catsListsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return itemId, tx.Commit()
}

func (r *CatPostgres) GetAll(userId, listId int) ([]cat.Cat, error) {
	var items []cat.Cat
	query := fmt.Sprintf(`SELECT ti.id, ti.age, ti.name, ti.filename, ti.description FROM %s ti INNER JOIN %s li on li.cat_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`, catsTable, catsListsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *CatPostgres) GetById(userId, listId, itemId int) (cat.Cat, error) {
	var item cat.Cat
	query := fmt.Sprintf(`SELECT ti.id, ti.age, ti.name, ti.filename, ti.description FROM %s ti INNER JOIN %s li on li.cat_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.cat_id = $1 AND ul.user_id = $2`,
		catsTable, catsListsTable, usersListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}
	FileName, err := RPresignedGetObject(r.minioDb, strconv.Itoa(itemId)+strconv.Itoa(listId)+"bucket", item.FileName+".png", time.Second*24*60*60)
	if err != nil {
		return item, err
	}
	item.FileName = FileName.String()
	return item, nil
}

func (r *CatPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul
									WHERE ti.id = li.cat_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		catsTable, catsListsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	if err != nil {
		return err
	}
	return nil
}

func (r *CatPostgres) Update(userId, listId, itemId int, input cat.UpdateCatInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Age != nil {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argId))
		args = append(args, *input.Age)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.FileName != nil {
		setValues = append(setValues, fmt.Sprintf("filename=$%d", argId))
		args = append(args, *input.FileName)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
									WHERE ti.id = li.cat_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		catsTable, setQuery, catsListsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)
	if err == nil && input.FileName != nil {
		err = RFPutObject(r.minioDb, strconv.Itoa(itemId)+strconv.Itoa(listId)+"bucket", *input.FileName+".png", *input.FileName+".png", "image/png")
		if err != nil {
			return err
		}
	}
	return err
}
