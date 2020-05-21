package compiler

// Error is a helper type for creating constant errors.
type Error string

func (e Error) Error() string { return string(e) }

const (
	// ErrUnsupported indicates that something is not supported. What exactly is
	// unsupported, must be indicated by a wrapping error.
	ErrUnsupported Error = "unsupported"
)
