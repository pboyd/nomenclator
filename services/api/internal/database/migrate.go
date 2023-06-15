package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// file driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate go-migrate database migration from file location
func Migrate(config *Config, db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("unable to create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+config.MigrationsDir, "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to build database migration %v", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
