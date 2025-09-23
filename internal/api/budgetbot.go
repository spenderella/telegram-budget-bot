package api

import (
	"fmt"
	"log"
	"strings"

	"telegram-finance-bot/internal/configs"
	"telegram-finance-bot/internal/models"
	"telegram-finance-bot/internal/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BudgetBot struct {
	config         *configs.Config
	debug          bool
	api            *tgbotapi.BotAPI
	expenseService *services.ExpenseService
}

// Конструктор
func NewBudgetBot(config *configs.Config, expenseService *services.ExpenseService) (*BudgetBot, error) {
	botAPI, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		return nil, err
	}

	return &BudgetBot{
		config:         config,
		debug:          true,
		api:            botAPI,
		expenseService: expenseService,
	}, nil
}

func (bot *BudgetBot) Start() {
	bot.api.Debug = bot.debug
	log.Printf("Authorized on account %s", bot.api.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = bot.config.PollingTimeout

	updates := bot.api.GetUpdatesChan(u)

	log.Println("Bot is running. Send /start to test...")
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			bot.handleMessage(&update)
		}
	}
}

func (bot *BudgetBot) formatExpenses(expenses []models.Expense) string {
	if len(expenses) == 0 {
		return "📊 Expenses not found"
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("📊 Found %d expenses:\n\n", len(expenses)))

	total := 0.0

	for _, expense := range expenses {
		dateStr := expense.Date.Format("02.01.2006 15:04")

		line := fmt.Sprintf("Amount: %.2f %s Category: %s\n Date: %s\n\n",
			expense.Amount,
			expense.Currency,
			expense.Category,
			dateStr)

		builder.WriteString(line)
		total += expense.Amount //convert when diff currency
	}

	builder.WriteString(fmt.Sprintf("Total amount: %.2f", total))
	return builder.String()
}
