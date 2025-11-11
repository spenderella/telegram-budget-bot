package api

import (
	"strings"
	"telegram-finance-bot/internal/models"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *BudgetBot) commandAddExpense(update *tgbotapi.Update) {

	userTgID := update.Message.From.ID
	username := update.Message.From.UserName

	amount, category, err := parseAddExpenseCommand(update.Message.Text)
	if err != nil {
		bot.sendReply(update, err.Error())
		return
	}

	err = bot.expenseService.AddExpense(userTgID, username, amount, category)
	if err != nil {
		bot.sendReply(update, "Failed to save expense: "+err.Error())
		return
	}

	bot.sendReply(update, "Expense saved successfully!")

}

func (bot *BudgetBot) commandGetExpenses(update *tgbotapi.Update) {

	userTgID := update.Message.From.ID
	parts := strings.Fields(update.Message.Text)
	dateFrom := time.Time{} // 0001-01-01
	dateTo := time.Now().UTC()

	filter := models.ExpenseFilter{
		UserTgID: userTgID,
		Limit:    &bot.config.GetExpenseLimit,
		DateFrom: &dateFrom,
		DateTo:   &dateTo,
	}

	if len(parts) > 1 {
		period := parts[1] // "today", "week", "month"
		dateFrom, dateTo := bot.getPeriodDates(period)
		filter.DateFrom = &dateFrom
		filter.DateTo = &dateTo
	}

	expenses, err := bot.expenseService.GetExpenses(filter)
	if err != nil {
		bot.sendReply(update, "Error getting expenses: "+err.Error())
		return
	}

	text := bot.formatExpenses(expenses)
	bot.sendReply(update, text)

}

func (bot *BudgetBot) commandGetStatistics(update *tgbotapi.Update) {

	userTgID := update.Message.From.ID
	parts := strings.Fields(update.Message.Text)
	dateFrom := time.Time{} // 0001-01-01
	dateTo := time.Now().UTC()

	filter := models.ExpenseFilter{
		UserTgID: userTgID,
		Limit:    &bot.config.GetExpenseLimit,
		DateFrom: &dateFrom,
		DateTo:   &dateTo,
	}

	if len(parts) > 1 {
		period := parts[1] // "today", "week", "month"
		dateFrom, dateTo := bot.getPeriodDates(period)
		filter.DateFrom = &dateFrom
		filter.DateTo = &dateTo
	}
	// tbd new query to db to get stat

	stat, err := bot.expenseService.GetStat(filter)
	if err != nil {
		bot.sendReply(update, err.Error())
		return
	}

	text := bot.formatCategoriesStat(stat)
	bot.sendReply(update, text)

}
