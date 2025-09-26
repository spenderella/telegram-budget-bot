package api

import (
	"telegram-finance-bot/internal/models"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *BudgetBot) commandAddExpense(update *tgbotapi.Update) {

	userTime := time.Unix(int64(update.Message.Date), 0)
	userID := update.Message.From.ID

	amount, category, err := parseAddExpenseCommand(update.Message.Text)
	if err != nil {
		bot.sendReply(update, err.Error())
		return
	}

	err = bot.expenseService.AddExpense(userID, amount, category, userTime)
	if err != nil {
		bot.sendReply(update, "Failed to save expense. Please try again later.")
		return
	}

	bot.sendReply(update, "Expense saved successfully!")

}

func (bot *BudgetBot) commandGetExpenses(update *tgbotapi.Update) {

	userID := update.Message.From.ID
	filter := models.ExpenseFilter{
		UserID: userID,
		Limit:  &bot.config.GetExpenseLimit,
	}

	expenses, err := bot.expenseService.GetExpenses(filter)
	if err != nil {
		bot.sendReply(update, "Error getting expenses: "+err.Error())
		return
	}

	text := bot.formatExpenses(expenses)
	bot.sendReply(update, text)

}
