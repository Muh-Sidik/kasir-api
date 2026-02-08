package service

import (
	"github.com/Muh-Sidik/kasir-api/internal/model"
	"github.com/Muh-Sidik/kasir-api/internal/model/dto"
	"github.com/Muh-Sidik/kasir-api/internal/repository"
	"github.com/gofrs/uuid/v5"
)

type ProductService interface {
	GetProduct(dto *dto.ProductQuery) ([]*model.ProductCategory, int, error)
	GetProductByID(id string) (*model.ProductCategory, error)
	CreateProduct(body *dto.ProductRequest) (*model.Product, error)
	UpdateProductByID(id string, body *dto.ProductRequest) (*model.Product, error)
	DeleteProductByID(id string) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) GetProduct(dto *dto.ProductQuery) ([]*model.ProductCategory, int, error) {
	return s.repo.GetProduct(dto)
}

func (s *productService) CreateProduct(body *dto.ProductRequest) (*model.Product, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	validCategory, err := uuid.FromString(body.CategoryID)
	if err != nil {
		return nil, err
	}

	return s.repo.CreateProduct(&model.Product{
		ID:         id,
		Name:       body.Name,
		Stock:      body.Stock,
		Price:      body.Price,
		CategoryID: validCategory,
	})
}

func (s *productService) GetProductByID(id string) (*model.ProductCategory, error) {
	return s.repo.GetProductByID(id)
}

func (s *productService) DeleteProductByID(id string) error {
	return s.repo.DeleteProductByID(id)
}

func (s *productService) UpdateProductByID(id string, body *dto.ProductRequest) (*model.Product, error) {
	validCategory, err := uuid.FromString(body.CategoryID)
	if err != nil {
		return nil, err
	}
	return s.repo.UpdateProductByID(id, &model.Product{
		Name:       body.Name,
		Stock:      body.Stock,
		Price:      body.Price,
		CategoryID: validCategory,
	})
}
