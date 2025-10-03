package api

import (
	"telegram-finance-bot/internal/models"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *BudgetBot) commandAddExpense(update *tgbotapi.Update) {

	userTime := time.Unix(int64(update.Message.Date), 0)
	userTgID := update.Message.From.ID
	username := update.Message.From.UserName

	amount, category, err := parseAddExpenseCommand(update.Message.Text)
	if err != nil {
		bot.sendReply(update, err.Error())
		return
	}

	err = bot.expenseService.AddExpense(userTgID, username, amount, category, userTime)
	if err != nil {
		bot.sendReply(update, "Failed to save expense: "+err.Error())
		return
	}

	bot.sendReply(update, "Expense saved successfully!")

}

func (bot *BudgetBot) commandGetExpenses(update *tgbotapi.Update) {

	userTgID := update.Message.From.ID
	filter := models.ExpenseFilter{
		UserID: userTgID,
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
