// Package database contains functions for interacting with the database.
package database

//go:generate sqlc generate

import (
	"database/sql"
	"fmt"

	// For postgres.
	_ "github.com/lib/pq"
)

// Config contains connection information for the database.
type Config struct {
	DSN           string
	MigrationsDir string
}

// Connect creates a new database connection.
func Connect(config *Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DSN)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	return db, nil
}
