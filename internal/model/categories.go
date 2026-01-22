package model

type Categories struct {
	ID          int    `sql:"id" json:"id"`
	Name        string `sql:"name" json:"name"`
	Description string `sql:"description" json:"description"`
}
