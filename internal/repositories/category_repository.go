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

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) GetCategories() ([]models.Category, error) {

	query := "SELECT id, name FROM categories LIMIT 100"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var cat models.Category

		err := rows.Scan(
			&cat.ID,
			&cat.Name,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, cat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
