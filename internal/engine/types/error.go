package types

import "fmt"

// Error is a sentinel error.
type Error string

func (e Error) Error() string { return string(e) }

// ErrTypeMismatch returns an error that indicates a type mismatch, and includes
// the expected and the actual type.
func ErrTypeMismatch(expected, got Type) Error {
	return Error(fmt.Sprintf("type mismatch: want %v, got %v", expected, got))
}
