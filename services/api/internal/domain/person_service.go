package domain

import (
	"context"
	"database/sql"
	"errors"
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
	if err := ps.validate(ctx, person); err != nil {
		return err
	}

	params := database.InsertPersonParams{
		Prefix: database.NullPrefix{
			Valid:  true,
			Prefix: database.Prefix(person.Prefix),
		},
		FirstName:  person.FirstName,
		MiddleName: sql.NullString{Valid: true, String: person.MiddleName},
		LastName:   person.LastName,
		Suffix: database.NullSuffix{
			Valid:  true,
			Suffix: database.Suffix(person.Suffix),
		},
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

func (ps *PersonService) validate(ctx context.Context, person *Person) error {
	err := ErrValidationFailed{}

	if person.Prefix != "" && !ps.queries.IsValidPrefix(ctx, person.Prefix) {
		err["prefix"] = "invalid prefix"
	}

	if person.Suffix != "" && !ps.queries.IsValidSuffix(ctx, person.Suffix) {
		err["suffix"] = "invalid suffix"
	}

	if len(err) > 0 {
		return err
	}

	return nil
}

// Load fetches a person by ID.
//
// If the person is not found, Load returns a nil person and a nil error.
func (ps *PersonService) Load(ctx context.Context, id int64) (*Person, error) {
	person, err := ps.queries.GetPerson(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to load person: %w", err)
	}

	return &Person{
		ID:         person.ID,
		Prefix:     string(person.Prefix.Prefix),
		FirstName:  person.FirstName,
		MiddleName: person.MiddleName.String,
		LastName:   person.LastName,
		Suffix:     string(person.Suffix.Suffix),
		CreatedAt:  person.CreatedAt,
		UpdatedAt:  person.UpdatedAt,
	}, nil
}
