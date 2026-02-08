package model

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Transaction struct {
	ID          uuid.UUID           `sql:"id" json:"id"`
	TotalAmount int64               `sql:"total_amount" json:"total_amount"`
	CreatedAt   time.Time           `sql:"created_at" json:"created_at"`
	Details     []TransactionDetail `json:"details"`
}

type TransactionDetail struct {
	ID            uuid.UUID `sql:"id" json:"id"`
	TransactionID uuid.UUID `sql:"transaction_id" json:"transaction_id"`
	ProductID     uuid.UUID `sql:"product_id" json:"product_id"`
	ProductName   string    `sql:"product_name,omitempty" json:"product_name,omitempty"`
	Quantity      int       `sql:"quantity" json:"quantity"`
	Subtotal      int64     `sql:"subtotal" json:"subtotal"`
	CreatedAt     time.Time `sql:"created_at" json:"created_at"`
}
