package model

import "time"

type Product struct {
	ID        int       `sql:"id" json:"id"`
	Name      string    `sql:"name" json:"name"`
	Price     int       `sql:"price" json:"price"`
	Stock     int       `sql:"stock" json:"stock"`
	CreatedAt time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt time.Time `sql:"updated_at" json:"updated_at"`
}
