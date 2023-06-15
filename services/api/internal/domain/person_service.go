package domain

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pboyd/nomenclator/api/internal/database"
)

// PersonService contains business logic for person related models.
type PersonService struct {
	queries *database.Queries
}

// NewPersonService creates a PersonService instance.
func NewPersonService(queries *database.Queries) *PersonService {
	return &PersonService{
		queries: queries,
	}
}

// Create creates a new person entry. ID, CreatedAt, and UpdatedAt are set
// on the passed in Person.
func (ps *PersonService) Create(ctx context.Context, person *Person) error {
	params := database.InsertPersonParams{
		Prefix:     sql.NullString{Valid: true, String: person.Prefix},
		FirstName:  person.FirstName,
		MiddleName: sql.NullString{Valid: true, String: person.MiddleName},
		LastName:   person.LastName,
		Suffix:     sql.NullString{Valid: true, String: person.Suffix},
	}

	if person.Prefix == "" {
		params.Prefix.Valid = false
	}
	if person.MiddleName == "" {
		params.MiddleName.Valid = false
	}
	if person.Suffix == "" {
		params.Suffix.Valid = false
	}

	result, err := ps.queries.InsertPerson(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to create person: %w", err)
	}

	person.ID = result.ID
	person.CreatedAt = result.CreatedAt
	person.UpdatedAt = result.UpdatedAt

	return nil
}
