package model

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Categories struct {
	ID          uuid.UUID       `sql:"id" json:"id"`
	Name        string    `sql:"name" json:"name"`
	Description string    `sql:"description" json:"description"`
	CreatedAt   time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt   time.Time `sql:"updated_at" json:"updated_at"`
}
