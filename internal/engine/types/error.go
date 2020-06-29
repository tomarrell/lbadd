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

// ErrCannotCast returns an error that indicates, that a case from the given
// from type to the given to type cannot be performed.
func ErrCannotCast(from, to Type) Error {
	return Error(fmt.Sprintf("cannot cast %v to %v", from, to))
}

// ErrDataSizeMismatch returns an error that indicates, that data which had an
// unexpected size was passed in. This will be useful for functions that expect
// fixed-size data.
func ErrDataSizeMismatch(expectedSize, gotSize int) Error {
	return Error(fmt.Sprintf("unexpected data size %v, need %v", gotSize, expectedSize))
}
