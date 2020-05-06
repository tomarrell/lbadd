package network

// Error is a helper type for creating constant errors.
type Error string

func (e Error) Error() string { return string(e) }

const (
	// ErrOpen indicates, that the component was already opened, and it is
	// unable to be opened another time.
	ErrOpen Error = "already open"
	// ErrClosed indicates, that the component is already closed, and it cannot
	// be used anymore.
	ErrClosed Error = "already closed"
)
