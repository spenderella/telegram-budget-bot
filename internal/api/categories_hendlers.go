package api

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *BudgetBot) commandGetCategories(update *tgbotapi.Update) {

	categories, err := bot.expenseService.CategoryService.GetCategories()
	if err != nil {
		bot.sendReply(update, "Error getting categories: "+err.Error())
		return
	}

	var builder strings.Builder
	builder.WriteString("Categories:\n\n")

	for _, category := range categories {
		builder.WriteString(fmt.Sprintf("%s\n", category.Name))
	}

	text := builder.String()
	bot.sendReply(update, text)

}
