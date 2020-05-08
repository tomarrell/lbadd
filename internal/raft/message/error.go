package message

// Error is a helper type that allows constant errors.
type Error string

func (e Error) Error() string { return string(e) }

const (
	// ErrUnknownKind indicates that the kind of the message is not known to
	// this implementation.
	ErrUnknownKind Error = "unknown message kind"
)
