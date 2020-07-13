package storage

import "fmt"

// Error is a sentinel error.
type Error string

func (e Error) Error() string { return string(e) }

// Sentinel errors.
const (
	ErrClosed          Error = "already closed"
	ErrNoSuchConfigKey Error = "no such configuration key"
)

// ErrNoSuchCell returns an error that indicates, that a cell with the given
// name could not be found.
func ErrNoSuchCell(cellKey string) Error {
	return Error(fmt.Sprintf("no such cell '%v'", cellKey))
}
