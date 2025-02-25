package cat

import "errors"

type CatList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type Cat struct {
	Id          int    `json:"id" db:"id"`
	Age         int    `json:"age" db:"age"`
	Name        string `json:"name" db:"name" binding:"required"`
	FileName    string `jsom:"filename" db:"filename"`
	Description string `json:"description" db:"description"`
}

type ListsItem struct {
	Id     int
	ListId int
	CatId  int
}

type UpdeteListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdeteListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateCatInput struct {
	Age         *int    `json:"age"`
	Name        *string `json:"name"`
	FileName    *string `jsom:"filename"`
	Description *string `json:"description"`
}

func (i UpdateCatInput) Validate() error {
	if i.Age == nil && i.Name == nil && i.FileName == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
