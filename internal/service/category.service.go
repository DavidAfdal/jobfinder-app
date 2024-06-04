package service

import (
	"github.com/DavidAfdal/workfinder/internal/entity"
	"github.com/DavidAfdal/workfinder/internal/repository"
	"github.com/google/uuid"
)


type CategoryService interface {
	FindAllCategory() ([]entity.Category, error)
	FindCategoryByID(id uuid.UUID) (*entity.Category, error)
	CreateCategory(category *entity.Category) (*entity.Category, error)
	UpdateCategory(category *entity.Category) (*entity.Category, error)
	DeleteCategory(id uuid.UUID) (bool, error)
}
type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo}
}

func (s *categoryService) FindAllCategory() ([]entity.Category, error) {
	return s.categoryRepo.FindAllCategory()
}

func (s *categoryService) FindCategoryByID(id uuid.UUID) (*entity.Category, error) {
	return s.categoryRepo.FindCategoryByID(id)
}

func (s *categoryService) CreateCategory(category *entity.Category) (*entity.Category, error) {
	return s.categoryRepo.CreateCategory(category)
}

func (s *categoryService) UpdateCategory(category *entity.Category) (*entity.Category, error) {
	return s.categoryRepo.UpdateCategory(category)
}

func (s *categoryService) DeleteCategory(id uuid.UUID) (bool, error) {
	category, err := s.categoryRepo.FindCategoryByID(id)

	if err != nil {
		return false, err
	}

	return s.categoryRepo.DeleteCategory(category)
}


