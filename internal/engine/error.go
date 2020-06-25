package engine

import "fmt"

// Error is a sentinel error.
type Error string

func (e Error) Error() string { return string(e) }

// Sentinel errors.
const (
	ErrClosed        Error = "already closed"
	ErrUnsupported   Error = "unsupported"
	ErrUnimplemented Error = "unimplemented"
)

// ErrNoSuchFunction returns an error indicating that an error with the given
// name can not be found.
func ErrNoSuchFunction(name string) error {
	return fmt.Errorf("no function for name %v(...)", name)
}
