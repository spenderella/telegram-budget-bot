# Telegram Budget Bot

Personal finance tracking bot for Telegram built with Go.

## Features

- Add expense: amount and category
- Get your expenses with optional filter by time period: today, week or month
- Get all existing categories
- Get expenses statistics by categories
- Automatic database migrations
- PostgreSQL database with Docker support

## Prerequisites

- Go 1.21+
- Docker and Docker Compose
- Telegram Bot Token (from @BotFather)
- PostgreSQL 15+ (via Docker)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/spenderella/telegram-budget-bot.git
cd telegram-budget-bot 
```
2. Create `.env` file with your configuration:
```bash
BOT_TOKEN=your_bot_token_here
DB_HOST=localhost
DB_PORT=5432
DB_USER=bot_user
DB_PASSWORD=bot_password
DB_NAME=budget_bot_db
DB_SSLMODE=disable
```

3. Start PostgreSQL database with Docker:
```bash
docker-compose up -d
```

Verify the database is running:
```bash
docker-compose ps
```
4. Run the bot:
```bash
go run cmd/budgetbot/main.go
```

The bot will automatically run database migrations on startup.

To stop the database:
```bash
docker-compose down
```

## Database

The bot uses PostgreSQL for storing expenses, categories, and users. Database migrations run automatically on application startup.

**Environment Variables:**
- `BOT_TOKEN` - Telegram Bot API token (from @BotFather)
- `DB_HOST` - PostgreSQL host (default: localhost)
- `DB_PORT` - PostgreSQL port (default: 5432)
- `DB_USER` - PostgreSQL username
- `DB_PASSWORD` - PostgreSQL password
- `DB_NAME` - Database name
- `DB_SSLMODE` - SSL mode (use "disable" for local development)

**Database Schema:**
- `users` - Telegram users
- `categories` - Expense categories (7 default categories)
- `expenses` - User expenses
- `schema_migrations` - Migration version tracking

**Default Categories:**
- food
- transportation
- house
- health
- entertainment
- personal
- other


## Usage

Available commands:

- `/start` - Welcome message
- `/help` - Show available commands
- `/add_expense <amount> <category>` - Add new expense with amount and category
- `/get_expenses [period]` - Get your expenses (limit 50). Optional period: today, week, month
- `/get_categories` - Get all expense categories
- `/get_statistics [period]` - Get spending totals by category (limit 50). Optional period: today, week, month

**Examples:**
```
/add_expense 25.50 coffee
/add_expense 100 food
/get_expenses today
/get_statistics month
```

## Development

### Project Structure

```
.
├── cmd/budgetbot/          # Application entry point
├── internal/
│   ├── api/                # Telegram bot handlers and parsers
│   ├── constants/          # SQL queries and constants
│   ├── database/           # Database connection and   migrations
│   ├── errors/             # Custom error definitions
│   ├── models/             # Data models
│   ├── repositories/       # Database access layer
│   └── services/           # Business logic layer
├── docker-compose.yaml     # PostgreSQL for development
├── docker-compose.test.yml # PostgreSQL for testing
└── .env                    # Environment configuration
```

### Running Tests

**Unit Tests:**
```bash
# Run all unit tests
go test -v ./internal/api/... ./internal/services/...

# Run specific package tests
go test -v ./internal/api
go test -v ./internal/services
```

**Integration Tests:**
```bash
# Start test database
docker-compose -f docker-compose.test.yml up -d

# Run all integration tests
go test -v ./internal/repositories

# Run specific repository tests
go test -v ./internal/repositories -run TestUserRepositoryTestSuite
go test -v ./internal/repositories -run TestCategoryRepositoryTestSuite
go test -v ./internal/repositories -run TestExpenseRepositoryTestSuite
```

**Test Coverage:**
- 33 unit tests (parsers, services)
- 17 integration tests (repositories)
- Total: 50 tests

### Architecture

The application follows a layered architecture:

1. **API Layer** (`internal/api`) - Telegram bot handlers, message parsing
2. **Service Layer** (`internal/services`) - Business logic, orchestration
3. **Repository Layer** (`internal/repositories`) - Database operations
4. **Models** (`internal/models`) - Data structures

All layers are covered by tests using testify/assert, testify/require, and gomock for mocking.


