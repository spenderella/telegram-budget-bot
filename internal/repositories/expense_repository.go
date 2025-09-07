package repositories

import "telegram-finance-bot/internal/models"

type ExpenseRepository struct {
	expenses map[int][]models.Expense // userID -> список трат
}
