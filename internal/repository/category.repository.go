package repository

import (
	"github.com/DavidAfdal/workfinder/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)



type CategoryRepository interface {
	CreateCategory(category *entity.Category) (*entity.Category, error)
	FindAllCategory() ([]entity.Category, error)
	FindCategoryByID(id uuid.UUID) (*entity.Category, error)
	UpdateCategory(category *entity.Category) (*entity.Category, error)
	DeleteCategory(category *entity.Category) (bool, error)
}

type categoryRepository struct {
	db *gorm.DB
}


func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}


func (r *categoryRepository) CreateCategory(category *entity.Category) (*entity.Category, error) {
	if err := r.db.Create(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (r *categoryRepository) FindAllCategory() ([]entity.Category, error) {
	categories := make([]entity.Category, 0)

	if err := r.db.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

func (r *categoryRepository) FindCategoryByID(id uuid.UUID) (*entity.Category, error) {
	category := new(entity.Category)

	if err := r.db.Where("id = ?", id).Preload("Jobs").First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (r *categoryRepository) UpdateCategory(category *entity.Category) (*entity.Category, error) {
	fields := make(map[string]interface{})

	if category.Title != "" {
		fields["title"] = category.Title
	}
	if category.Icon != "" {
		fields["icon"] = category.Icon
	}

	if err := r.db.Model(&category).Updates(fields).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (r *categoryRepository) DeleteCategory(category *entity.Category) (bool, error) {
	if err := r.db.Delete(&category).Error; err != nil {
		return false, err
	}
	return true, nil
}
