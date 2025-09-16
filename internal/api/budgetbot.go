package api

import (
	"log"

	"telegram-finance-bot/internal/configs"
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
