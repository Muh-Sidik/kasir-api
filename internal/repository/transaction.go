package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Muh-Sidik/kasir-api/internal/model"
	"github.com/Muh-Sidik/kasir-api/internal/model/dto"
	"github.com/gofrs/uuid/v5"
)

type TransactionRepository interface {
	CreateTransaction(items []dto.CheckoutItem) (*model.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (t *transactionRepository) CreateTransaction(items []dto.CheckoutItem) (*model.Transaction, error) {
	if len(items) == 0 {
		return nil, errors.New("items cannot be empty")
	}

	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	producttock, err := t.validateAndStockLock(tx, items)
	if err != nil {
		return nil, err
	}

	transactionID, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("failed generate id: %w", err)
	}
	totalAmount := int64(0)
	details := make([]model.TransactionDetail, 0, len(items))
	for i, item := range items {
		subTotal := int64(producttock[i].Price) * int64(item.Quantity)
		totalAmount += subTotal

		_, err = tx.Exec(
			"UPDATE product SET stock = stock - $1 WHERE id = $2",
			item.Quantity,
			item.ProductID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to update stock for product %s: %w", item.ProductID, err)
		}
		detailID, err := uuid.NewV7()
		if err != nil {
			return nil, fmt.Errorf("generate detail id failed: %w", err)
		}

		details = append(details, model.TransactionDetail{
			ID:            detailID,
			ProductID:     item.ProductID,
			TransactionID: transactionID,
			ProductName:   producttock[i].Name,
			Quantity:      item.Quantity,
			Subtotal:      subTotal,
		})
	}

	_, err = tx.Exec("INSERT INTO transactions (id, total_amount, created_at) VALUES ($1,$2,NOW())", transactionID, totalAmount)
	if err != nil {
		return nil, err
	}

	if err := t.bulkInsertDetails(tx, details); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &model.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}

func (t *transactionRepository) validateAndStockLock(tx *sql.Tx, items []dto.CheckoutItem) ([]model.Product, error) {
	productIDs := make([]any, len(items))
	placeholders := make([]string, len(items))
	for i, item := range items {
		productIDs[i] = item.ProductID
		// Query dengan FOR UPDATE untuk locking row
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf(
		`SELECT id, name, price, stock
         FROM product
         WHERE id IN (%s)
         FOR UPDATE`,
		strings.Join(placeholders, ","),
	)

	rows, err := tx.Query(query, productIDs...)
	if err != nil {
		return nil, fmt.Errorf("query product failed: %w", err)
	}
	defer rows.Close()

	// ✅ GUNAKAN MAP UNTUK LOOKUP AMAN (tanpa asumsi urutan)
	productMap := make(map[uuid.UUID]struct {
		Name  string
		Price int
		Stock int
	})

	for rows.Next() {
		var id uuid.UUID
		var name string
		var price, stock int
		if err := rows.Scan(&id, &name, &price, &stock); err != nil {
			return nil, fmt.Errorf("scan product failed: %w", err)
		}
		productMap[id] = struct {
			Name  string
			Price int
			Stock int
		}{name, price, stock}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate product failed: %w", err)
	}

	results := make([]model.Product, len(items))
	for i, item := range items {
		prod, exists := productMap[item.ProductID]
		if !exists {
			return nil, fmt.Errorf("product %s not found", item.ProductID)
		}
		if prod.Stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for %s: have %d, need %d",
				item.ProductID, prod.Stock, item.Quantity)
		}
		results[i] = model.Product{
			ID:    item.ProductID,
			Name:  prod.Name,
			Price: prod.Price,
			Stock: prod.Stock,
		}
	}

	return results, nil
}

func (t *transactionRepository) bulkInsertDetails(tx *sql.Tx, details []model.TransactionDetail) error {
	if len(details) == 0 {
		return nil
	}

	valueStrings := make([]string, len(details))
	args := make([]any, 0, len(details)*6)
	for i, d := range details {
		args = append(args, d.ID, d.TransactionID, d.ProductID, d.Quantity, d.Subtotal)
		valueStrings[i] = fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,NOW())", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5)
	}

	query := fmt.Sprintf(
		"INSERT INTO transaction_details (id,transaction_id,product_id,quantity,subtotal,created_at) VALUES %s",
		strings.Join(valueStrings, ","),
	)

	_, err := tx.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("bulk insert details failed: %w", err)
	}
	return nil
}
