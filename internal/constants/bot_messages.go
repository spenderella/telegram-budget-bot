package constants

const (
	StartMessage = `Welcome to Budget Bot, %s!

Available commands:
/help - Show commands
/add_expense <amount> <category> - Add expense
/get_expenses - Get your expenses (limit 50). Optional: today, week, month

Examples: 
/add_expense 25.50 coffee
/get_expenses today`

	HelpMessage = `Commands:
/start - Welcome  
/help - Show this help

/add_expense <amount> <category> - Add expense
/get_expenses - Get your expenses (limit 50). Optional: today, week, month

Examples: 
/add_expense 25.50 coffee
/get_expenses today`
)
