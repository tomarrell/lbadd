package engine

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/engine/types"
)

// Error is a sentinel error.
type Error string

func (e Error) Error() string { return string(e) }

const (
	// ErrClosed indicates that the component can not be used anymore, because
	// it already has been closed.
	ErrClosed Error = "already closed"
	// ErrUnsupported indicates that a requested feature is explicitely not
	// supported. This is different from ErrUnimplemented, since
	// ErrUnimplemented indicates, that the feature has not been implemented
	// yet, while ErrUnsupported indicates, that the feature is intentionally
	// unimplemented.
	ErrUnsupported Error = "unsupported"
	// ErrUnimplemented indicates a missing implementation for the requested
	// feature. It may be implemented in the next version.
	ErrUnimplemented Error = "unimplemented"
)

// ErrNoSuchFunction returns an error indicating that an error with the given
// name can not be found.
func ErrNoSuchFunction(name string) error {
	return fmt.Errorf("no function for name %v(...)", name)
}

// ErrUncomparable returns an error indicating that the given type does not
// implement the types.Comparator interface, and thus, values of that type
// cannot be compared.
func ErrUncomparable(t types.Type) error {
	return fmt.Errorf("type %v is not comparable", t)
}
