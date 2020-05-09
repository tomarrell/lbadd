package cluster

// Error is a helper type for creating constant errors.
type Error string

func (e Error) Error() string { return string(e) }

const (
	// ErrTimeout indicates, that a the operation took longer than allowed.
	// Maybe there was a deadline from a context.
	ErrTimeout Error = "timeout"
)
