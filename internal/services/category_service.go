package services

import (
	//"database/sql"

	"telegram-finance-bot/internal/models"
	//"telegram-finance-bot/internal/repositories"
)

type ICategoryRepository interface {
	GetCategory(name string) (*models.Category, error)
	GetCategories() ([]models.Category, error)
}

type CategoryService struct {
	categoryRepo ICategoryRepository
}

func NewCategoryService(categoryRepo ICategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (s *CategoryService) GetCategory(name string) (*models.Category, error) {
	category, err := s.categoryRepo.GetCategory(name)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) GetCategories() ([]models.Category, error) {
	return s.categoryRepo.GetCategories()
}
