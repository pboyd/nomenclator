// Package domain application layer of service
package domain

import (
	"database/sql"

	"github.com/pboyd/nomenclator/api/internal/database"
)

// Bundle contains all the services for convenience.
type Bundle struct {
	PersonService *PersonService
}

// NewBundle create new service bundle
func NewBundle(db *sql.DB) *Bundle {
	q := database.New(db)

	return &Bundle{
		PersonService: NewPersonService(q),
	}
}
