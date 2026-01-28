package repository

import (
	"database/sql"

	"github.com/Muh-Sidik/kasir-api/internal/model"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/request"
)

type CategoryRepository interface {
	GetCategories(paginate *request.PaginateRes) ([]*model.Categories, int, error)
	GetCategoryByID(id int) (*model.Categories, error)
	CreateCategory(category *model.Categories) (*model.Categories, error)
	UpdateCategoryByID(id int, category *model.Categories) (*model.Categories, error)
	DeleteCategoryByID(id int) error
}

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) GetCategories(paginate *request.PaginateRes) ([]*model.Categories, int, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM categories WHERE 1=1 LIMIT ? OFFSET ?`

	rows, err := c.db.Query(query, paginate.Limit, paginate.Offset)

	if err != nil {
		return nil, 0, err
	}

	if rows.Err() != nil {
		return nil, 0, rows.Err()
	}

	defer rows.Close()

	categories := make([]*model.Categories, 0)

	for rows.Next() {
		var category model.Categories
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
			&category.UpdatedAt,
		)

		if err != nil {
			return nil, 0, err
		}

		categories = append(categories, &category)
	}

	var total int
	err = c.db.QueryRow(`SELECT COUNT(*) FROM categories`).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

func (c *categoryRepo) CreateCategory(body *model.Categories) (*model.Categories, error) {
	rows := c.db.QueryRow(
		`INSERT INTO categories(id, name, description, created_at, updated_at) VALUES(?,?,NOW(),NOW()) RETURNING id, created_at, updated_at`,
		body.ID,
		body.Name,
		body.Description,
	)

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var category model.Categories
	if err := rows.Scan(
		&category.ID,
		&category.CreatedAt,
		&category.UpdatedAt,
	); err != nil {
		return nil, err
	}

	category.Name = body.Name
	category.Description = body.Description

	return &category, nil
}

func (c *categoryRepo) GetCategoryByID(id int) (*model.Categories, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM categories WHERE id = ?`

	rows := c.db.QueryRow(query, id)

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var category model.Categories
	err := rows.Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *categoryRepo) DeleteCategoryByID(id int) error {
	_, err := c.db.Exec(
		`DELETE FROM categories WHERE id = ?`,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepo) UpdateCategoryByID(id int, body *model.Categories) (*model.Categories, error) {
	rows := c.db.QueryRow(
		`UPDATE categories SET name = ?, description = ?, updated_at = NOW() WHERE id = ? RETURNING id, created_at, updated_at`,
		body.Name,
		body.Description,
		id,
	)

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var category model.Categories
	if err := rows.Scan(
		&category.ID,
		&category.CreatedAt,
		&category.UpdatedAt,
	); err != nil {
		return nil, err
	}

	category.Name = body.Name
	category.Description = body.Description

	return &category, nil
}
