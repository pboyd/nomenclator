// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: people.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const getPerson = `-- name: GetPerson :one
SELECT id, created_at, updated_at, prefix, first_name, middle_name, last_name, suffix FROM people WHERE id = $1
`

// Get one person by ID.
func (q *Queries) GetPerson(ctx context.Context, id int64) (Person, error) {
	row := q.db.QueryRowContext(ctx, getPerson, id)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Prefix,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.Suffix,
	)
	return i, err
}

const insertPerson = `-- name: InsertPerson :one
INSERT INTO people (
    prefix,
    first_name,
    middle_name,
    last_name,
    suffix
) VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at
`

type InsertPersonParams struct {
	Prefix     NullPrefix
	FirstName  string
	MiddleName sql.NullString
	LastName   string
	Suffix     NullSuffix
}

type InsertPersonRow struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Add one person.
func (q *Queries) InsertPerson(ctx context.Context, arg InsertPersonParams) (InsertPersonRow, error) {
	row := q.db.QueryRowContext(ctx, insertPerson,
		arg.Prefix,
		arg.FirstName,
		arg.MiddleName,
		arg.LastName,
		arg.Suffix,
	)
	var i InsertPersonRow
	err := row.Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	return i, err
}

const listPrefixes = `-- name: ListPrefixes :many
SELECT unnest(enum_range(null::prefix))::text prefix
`

func (q *Queries) ListPrefixes(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listPrefixes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var prefix string
		if err := rows.Scan(&prefix); err != nil {
			return nil, err
		}
		items = append(items, prefix)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSuffixes = `-- name: ListSuffixes :many
SELECT unnest(enum_range(null::suffix))::text suffix
`

func (q *Queries) ListSuffixes(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listSuffixes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var suffix string
		if err := rows.Scan(&suffix); err != nil {
			return nil, err
		}
		items = append(items, suffix)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
