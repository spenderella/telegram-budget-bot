package api

import (
	"fmt"
	"log"
	"strings"

	"telegram-finance-bot/internal/configs"
	"telegram-finance-bot/internal/models"

	//"telegram-finance-bot/internal/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ExpenseService interface {
	AddExpense(telegramID int64, username string, amount float64, categoryName string) error
	GetExpenses(filter models.ExpenseFilter) ([]models.Expense, error)
	GetStat(filter models.ExpenseFilter) ([]models.CategoryExpenses, error)
}

type CategoryService interface {
	GetCategories() ([]models.Category, error)
}

type BudgetBot struct {
	config          *configs.Config
	debug           bool
	api             *tgbotapi.BotAPI
	expenseService  ExpenseService
	categoryService CategoryService
}

func NewBudgetBot(config *configs.Config, expenseService ExpenseService, categoryService CategoryService) (*BudgetBot, error) {
	botAPI, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		return nil, err
	}

	return &BudgetBot{
		config:          config,
		debug:           true,
		api:             botAPI,
		expenseService:  expenseService,
		categoryService: categoryService,
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
			expense.Category.Name,
			dateStr)

		builder.WriteString(line)
		total += expense.Amount //convert when diff currency
	}

	builder.WriteString(fmt.Sprintf("Total amount: %.2f", total))
	return builder.String()
}

func (bot *BudgetBot) formatCategoriesStat(stat []models.CategoryExpenses) string {
	if len(stat) == 0 {
		return "📊 Data not found"
	}

	var builder strings.Builder
	builder.WriteString("📊 Statistics for categories:\n\n")

	for _, catStat := range stat {

		builder.WriteString(fmt.Sprintf("Category: %s   Total: %.2f %s\n",
			catStat.Category.Name,
			catStat.Total,
			catStat.TotalCurrency))
	}

	return builder.String()
}
