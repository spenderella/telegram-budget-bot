package main

import (
	"log"

	"github.com/joho/godotenv"

	"telegram-finance-bot/internal/api"
	"telegram-finance-bot/internal/configs"
	"telegram-finance-bot/internal/database"
	"telegram-finance-bot/internal/repositories"
	"telegram-finance-bot/internal/services"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config, err := configs.LoadConfig("internal/configs/configs.json")
	if err != nil {
		log.Fatal("Failed to read configs:", err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	categoryRepo := repositories.NewCategoryRepository(db)
	userRepo := repositories.NewUserRepository(db)
	expenseRepo := repositories.NewExpenseRepository(db)

	categoryService := services.NewCategoryService(categoryRepo)
	userService := services.NewUserService(userRepo)
	expenseService := services.NewExpenseService(expenseRepo, userService, categoryService)

	bot, err := api.NewBudgetBot(config, expenseService, categoryService)
	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}

	bot.Start()
}
