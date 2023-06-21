-- name: InsertPerson :one
-- Add one person.
INSERT INTO people (
    prefix,
    first_name,
    middle_name,
    last_name,
    suffix
) VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at;

-- name: GetPerson :one
-- Get one person by ID.
SELECT * FROM people WHERE id = $1;

-- name: ListPrefixes :many
SELECT unnest(enum_range(null::prefix))::text prefix;

-- name: ListSuffixes :many
SELECT unnest(enum_range(null::suffix))::text suffix;
