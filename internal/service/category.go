package service

import (
	"github.com/Muh-Sidik/kasir-api/internal/model"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/request"
	"github.com/Muh-Sidik/kasir-api/internal/repository"
)

type CategoryService interface {
	GetCategories(paginate *request.PaginateRes) ([]*model.Categories, int, error)
	GetCategoryByID(id int) (*model.Categories, error)
	CreateCategory(category *model.Categories) (*model.Categories, error)
	UpdateCategoryByID(id int, category *model.Categories) (*model.Categories, error)
	DeleteCategoryByID(id int) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) GetCategories(paginate *request.PaginateRes) ([]*model.Categories, int, error) {
	return s.categoryRepo.GetCategories(paginate)
}

func (s *categoryService) CreateCategory(category *model.Categories) (*model.Categories, error) {
	return s.categoryRepo.CreateCategory(category)
}

func (s *categoryService) GetCategoryByID(id int) (*model.Categories, error) {
	return s.categoryRepo.GetCategoryByID(id)
}

func (s *categoryService) DeleteCategoryByID(id int) error {
	return s.categoryRepo.DeleteCategoryByID(id)
}

func (s *categoryService) UpdateCategoryByID(id int, category *model.Categories) (*model.Categories, error) {
	return s.categoryRepo.UpdateCategoryByID(id, category)
}
