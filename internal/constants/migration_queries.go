package constants

const (
	CreateUsersTable = `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        telegram_id BIGINT UNIQUE NOT NULL,
        username VARCHAR(255),
        created_at TIMESTAMPTZ DEFAULT NOW()
    );`

	CreateCategoriesTable = `
    CREATE TABLE IF NOT EXISTS categories (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) UNIQUE NOT NULL,
        created_at TIMESTAMPTZ DEFAULT NOW()
    );`

	CreateExpensesTable = `
    CREATE TABLE IF NOT EXISTS expenses (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE RESTRICT,
        amount DECIMAL(10, 2) NOT NULL CHECK (amount > 0),
        currency VARCHAR(3) NOT NULL,
        created_at TIMESTAMPTZ DEFAULT NOW()
    );`

	CreateMigrationsTable = `
    CREATE TABLE IF NOT EXISTS schema_migrations (
        version VARCHAR(255) PRIMARY KEY,
        applied_at TIMESTAMPTZ DEFAULT NOW()
    );`

	InsertDefaultCategories = `
    INSERT INTO categories (name) VALUES 
        ('food'),
        ('tranportasion'),
        ('house'),
        ('health'),
        ('entertaiment'),
        ('personal'),
        ('other')
    ON CONFLICT (name) DO NOTHING;
    `
)
