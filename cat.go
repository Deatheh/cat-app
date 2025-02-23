package cat

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
	Description string `json:"description" db:"description"`
}

type ListsItem struct {
	Id     int
	ListId int
	CatId  int
}

type Foto struct {
	Id  int    `json:"id"`
	Url string `json:"url" binding:"required"`
}

type CatFoto struct {
	Id    int
	CatId int
	Foto  int
}
