package repository

import (
	"database/sql"

	"github.com/Muh-Sidik/kasir-api/internal/model"
	"github.com/Muh-Sidik/kasir-api/internal/model/dto"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/request"
)

type ProductRepository interface {
	GetProduct(paginate *request.PaginateRes) ([]*dto.ProductCategory, int, error)
	CreateProduct(body *model.Product) (*model.Product, error)
	GetProductByID(id int) (*dto.ProductCategory, error)
	DeleteProductByID(id int) error
	UpdateProductByID(id int, body *model.Product) (*model.Product, error)
	GetProductByCategoryID(categoryID int, paginate *request.PaginateRes) ([]*dto.ProductCategory, int, error)
}

type productRepo struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) GetProduct(paginate *request.PaginateRes) ([]*dto.ProductCategory, int, error) {
	query := `SELECT 
		p.id,
		p.name, 
		p.price, 
		p.stock, 
		c.name as category_name,
		p.created_at, 
		p.updated_at 
	FROM product p
	WHERE 1=1
	JOIN categories c ON p.category_id = c.id
	LIMIT ? OFFSET ?`

	rows, err := p.db.Query(query, paginate.Limit, paginate.Offset)

	if err != nil {
		return nil, 0, err
	}

	if rows.Err() != nil {
		return nil, 0, rows.Err()
	}

	defer rows.Close()

	listProduct := make([]*dto.ProductCategory, 0)

	for rows.Next() {
		var product dto.ProductCategory
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.CategoryName,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, 0, err
		}

		listProduct = append(listProduct, &product)
	}

	var total int
	err = p.db.QueryRow(`
		SELECT COUNT(*) FROM product p
		JOIN categories c ON p.category_id = c.id
		WHERE 1=1
	`).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return listProduct, total, nil
}

func (p *productRepo) CreateProduct(body *model.Product) (*model.Product, error) {
	rows := p.db.QueryRow(
		`INSERT INTO product(id, name, price, stock, category_id, created_at, updated_at) VALUES(?,?,?, ?, NOW(), NOW()) RETURNING id, created_at, updated_at`,
		body.ID,
		body.Name,
		body.Price,
		body.Stock,
		body.CategoryID,
	)

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var product model.Product
	if err := rows.Scan(
		&product.ID,
		&product.CreatedAt,
		&product.UpdatedAt,
	); err != nil {
		return nil, err
	}

	product.Name = body.Name
	product.Price = body.Price
	product.Stock = body.Stock

	return &product, nil
}

func (p *productRepo) GetProductByID(id int) (*dto.ProductCategory, error) {
	query := `SELECT 
		p.id,
		p.name, 
		p.price, 
		p.stock, 
		c.name as category_name,
		p.created_at, 
		p.updated_at 
	FROM product p
	WHERE p.id = ? 
	JOIN categories c ON p.category_id = c.id`

	rows := p.db.QueryRow(query, id)

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var product dto.ProductCategory
	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.CategoryName,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productRepo) DeleteProductByID(id int) error {
	_, err := p.db.Exec(
		`DELETE FROM product WHERE id = ?`,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *productRepo) UpdateProductByID(id int, body *model.Product) (*model.Product, error) {
	rows := p.db.QueryRow(
		`UPDATE product SET name = ?, price = ?, stock = ?, category_id = ?, updated_at = NOW() WHERE id = ? RETURNING id, created_at, updated_at`,
		body.Name,
		body.Price,
		body.Stock,
		body.CategoryID,
		id,
	)

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var product model.Product
	if err := rows.Scan(
		&product.ID,
		&product.CreatedAt,
		&product.UpdatedAt,
	); err != nil {
		return nil, err
	}

	product.Name = body.Name
	product.Price = body.Price
	product.Stock = body.Stock

	return &product, nil
}

func (p *productRepo) GetProductByCategoryID(categoryID int, paginate *request.PaginateRes) ([]*dto.ProductCategory, int, error) {
	query := `SELECT 
		p.id,
		p.name, 
		p.price, 
		p.stock, 
		c.name as category_name,
		p.created_at, 
		p.updated_at 
	FROM product p
	WHERE p.category_id = ? 
	JOIN categories c ON p.category_id = c.id
	LIMIT ? OFFSET ?`

	rows, err := p.db.Query(query, categoryID, paginate.Limit, paginate.Offset)

	if err != nil {
		return nil, 0, err
	}

	if rows.Err() != nil {
		return nil, 0, rows.Err()
	}

	defer rows.Close()

	listProduct := make([]*dto.ProductCategory, 0)

	for rows.Next() {
		var product dto.ProductCategory
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.CategoryName,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, 0, err
		}

		listProduct = append(listProduct, &product)
	}

	var total int
	err = p.db.QueryRow(`
		SELECT COUNT(*) FROM product p
		JOIN categories c ON p.category_id = c.id
		WHERE p.category_id = ?
	`, categoryID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return listProduct, total, nil
}
