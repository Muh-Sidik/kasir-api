package dto

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type ProductCategory struct {
	ID           uuid.UUID `sql:"id" json:"id"`
	Name         string    `sql:"name" json:"name"`
	Price        int       `sql:"price" json:"price"`
	Stock        int       `sql:"stock" json:"stock"`
	CategoryName string    `sql:"category_name" json:"category_name"`
	CreatedAt    time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt    time.Time `sql:"updated_at" json:"updated_at"`
}
