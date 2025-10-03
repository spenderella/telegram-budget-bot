package repositories

import (
	"database/sql"
	//"log"
	"strings"

	"telegram-finance-bot/internal/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetCategory(name string) (*models.Category, error) {
	var category models.Category
	categoryName := strings.ToLower(name)
	query := "SELECT id, name FROM categories WHERE name = $1"

	err := r.db.QueryRow(query, categoryName).Scan(
		&category.ID, &category.Name,
	)

	return &category, err
}
