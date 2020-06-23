package engine

// Error is a sentinel error.
type Error string

func (e Error) Error() string { return string(e) }

// Sentinel errors.
const (
	ErrClosed      Error = "already closed"
	ErrUnsupported Error = "unsupported"
)
