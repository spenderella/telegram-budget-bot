package main

import (
	"log"
	"os"

	"telegram-finance-bot/internal/api"
	"telegram-finance-bot/internal/models"
	"telegram-finance-bot/internal/repositories"
	"telegram-finance-bot/internal/services"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is not set")
	}

	ER := repositories.ExpenseRepository{
		Expenses: make(map[int64][]models.Expense),
	}
	ES := services.ExpenseService{
		Repository: &ER,
	}

	bot, err := api.NewBudgetBot(token, &ES)
	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}

	bot.Start()
}
