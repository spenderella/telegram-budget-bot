package database

import (
	"database/sql"
	"log"

	"telegram-finance-bot/internal/constants"
)

func RunMigrations(db *sql.DB) error {
	if _, err := db.Exec(constants.CreateMigrationsTable); err != nil {
		return err
	}

	migrations := map[string]string{
		"001_create_users_table":      constants.CreateUsersTable,
		"002_create_categories_table": constants.CreateCategoriesTable,
		"003_create_expenses_table":   constants.CreateExpensesTable,
		"004_insert_categories":       constants.InsertDefaultCategories,
	}

	for version, sqlQuery := range migrations {
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM schema_migrations WHERE version = $1", version).Scan(&count)
		if err != nil {
			return err
		}

		if count == 0 {
			if _, err := db.Exec(sqlQuery); err != nil {
				return err
			}

			_, err = db.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", version)
			if err != nil {
				return err
			}

			log.Printf("Migration %s applied successfully", version)
		} else {
			log.Printf("Migration %s already applied, skipping", version)
		}
	}

	return nil
}
