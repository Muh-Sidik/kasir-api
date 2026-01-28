package service

import (
	"github.com/Muh-Sidik/kasir-api/internal/model"
	"github.com/Muh-Sidik/kasir-api/internal/model/dto"
	"github.com/Muh-Sidik/kasir-api/internal/model/dto/reqdto"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/request"
	"github.com/Muh-Sidik/kasir-api/internal/repository"
	"github.com/gofrs/uuid/v5"
)

type ProductService interface {
	GetProduct(paginate *request.PaginateRes) ([]*dto.ProductCategory, int, error)
	GetProductByID(id string) (*dto.ProductCategory, error)
	CreateProduct(body *reqdto.ProductRequest) (*model.Product, error)
	UpdateProductByID(id string, body *reqdto.ProductRequest) (*model.Product, error)
	DeleteProductByID(id string) error
	GetProductByCategoryID(categoryID string, paginate *request.PaginateRes) ([]*dto.ProductCategory, int, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) GetProduct(paginate *request.PaginateRes) ([]*dto.ProductCategory, int, error) {
	return s.repo.GetProduct(paginate)
}

func (s *productService) CreateProduct(body *reqdto.ProductRequest) (*model.Product, error) {
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

func (s *productService) GetProductByID(id string) (*dto.ProductCategory, error) {
	return s.repo.GetProductByID(id)
}

func (s *productService) DeleteProductByID(id string) error {
	return s.repo.DeleteProductByID(id)
}

func (s *productService) UpdateProductByID(id string, body *reqdto.ProductRequest) (*model.Product, error) {
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

func (s *productService) GetProductByCategoryID(categoryID string, paginate *request.PaginateRes) ([]*dto.ProductCategory, int, error) {
	return s.repo.GetProductByCategoryID(categoryID, paginate)
}
