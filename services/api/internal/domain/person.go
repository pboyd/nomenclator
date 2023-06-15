package domain

import "time"

// Person hold information about one person.
type Person struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Prefix     string `json:"prefix"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Suffix     string `json:"suffix"`
}
