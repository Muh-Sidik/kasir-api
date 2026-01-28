package model

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Product struct {
	ID         uuid.UUID       `sql:"id" json:"id"`
	Name       string    `sql:"name" json:"name"`
	Price      int       `sql:"price" json:"price"`
	Stock      int       `sql:"stock" json:"stock"`
	CategoryID uuid.UUID    `sql:"category_id" json:"category_id"`
	CreatedAt  time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at" json:"updated_at"`
}
