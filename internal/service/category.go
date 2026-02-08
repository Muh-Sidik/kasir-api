package service

import (
	"github.com/Muh-Sidik/kasir-api/internal/model"
	"github.com/Muh-Sidik/kasir-api/internal/model/dto"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/request"
	"github.com/Muh-Sidik/kasir-api/internal/repository"
	"github.com/gofrs/uuid/v5"
)

type CategoryService interface {
	GetCategories(paginate *request.PaginateQuery, search string) ([]*model.Categories, int, error)
	GetCategoryByID(id string) (*model.Categories, error)
	CreateCategory(category *dto.CategoryRequest) (*model.Categories, error)
	UpdateCategoryByID(id string, category *dto.CategoryRequest) (*model.Categories, error)
	DeleteCategoryByID(id string) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) GetCategories(paginate *request.PaginateQuery, search string) ([]*model.Categories, int, error) {
	return s.categoryRepo.GetCategories(paginate, search)
}

func (s *categoryService) CreateCategory(req *dto.CategoryRequest) (*model.Categories, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return s.categoryRepo.CreateCategory(&model.Categories{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	})
}

func (s *categoryService) GetCategoryByID(id string) (*model.Categories, error) {
	return s.categoryRepo.GetCategoryByID(id)
}

func (s *categoryService) DeleteCategoryByID(id string) error {
	return s.categoryRepo.DeleteCategoryByID(id)
}

func (s *categoryService) UpdateCategoryByID(id string, req *dto.CategoryRequest) (*model.Categories, error) {
	return s.categoryRepo.UpdateCategoryByID(id, &model.Categories{
		Name:        req.Name,
		Description: req.Description,
	})
}
