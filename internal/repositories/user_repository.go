package repositories

import (
	"database/sql"
	//"log"

	"telegram-finance-bot/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUser(telegramID int64) (*models.User, error) {
	var user models.User
	query := "SELECT id, telegram_id, username FROM users WHERE telegram_id = $1"

	err := r.db.QueryRow(query, telegramID).Scan(
		&user.ID, &user.TelegramID, &user.Username,
	)

	return &user, err
}

func (r *UserRepository) Create(telegramID int64, username string) (*models.User, error) {
	var user models.User
	query := `
        INSERT INTO users (telegram_id, username) 
        VALUES ($1, $2)
        RETURNING id, telegram_id, username
    `

	err := r.db.QueryRow(query, telegramID, username).Scan(
		&user.ID, &user.TelegramID, &user.Username)

	return &user, err
}
