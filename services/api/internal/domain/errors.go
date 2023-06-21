package domain

import "fmt"

// ErrValidationFailed is returned when one or more fields is invalid.
//
// It is a map of field names to error message.
type ErrValidationFailed map[string]string

func (e ErrValidationFailed) Error() string {
	return fmt.Sprintf("validation failed: %v", map[string]string(e))
}
