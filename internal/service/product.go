package service

import (
	"github.com/Muh-Sidik/kasir-api/internal/model"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/request"
	"github.com/Muh-Sidik/kasir-api/internal/repository"
)

type ProductService interface {
	GetProduct(paginate *request.PaginateRes) ([]*model.Product, int, error)
	GetProductByID(id int) (*model.Product, error)
	CreateProduct(body *model.Product) (*model.Product, error)
	UpdateProductByID(id int, body *model.Product) (*model.Product, error)
	DeleteProductByID(id int) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) GetProduct(paginate *request.PaginateRes) ([]*model.Product, int, error) {
	return s.repo.GetProduct(paginate)
}

func (s *productService) CreateProduct(body *model.Product) (*model.Product, error) {
	return s.repo.CreateProduct(body)
}

func (s *productService) GetProductByID(id int) (*model.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *productService) DeleteProductByID(id int) error {
	return s.repo.DeleteProductByID(id)
}

func (s *productService) UpdateProductByID(id int, body *model.Product) (*model.Product, error) {
	return s.repo.UpdateProductByID(id, body)
}
