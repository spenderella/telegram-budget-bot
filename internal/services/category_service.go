package services

import (
	//"database/sql"

	"telegram-finance-bot/internal/models"
	"telegram-finance-bot/internal/repositories"
)

type CategoryService struct {
	categoryRepo *repositories.CategoryRepository
}

func NewCategoryService(categoryRepo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (s *CategoryService) GetCategory(name string) (*models.Category, error) {
	category, err := s.categoryRepo.GetCategory(name)

	if err != nil {
		return nil, err
	}

	return category, nil
}
