package message

// Error is a helper type that allows constant errors.
type Error string

func (e Error) Error() string { return string(e) }

const (
	// ErrUnknownKind indicates that the kind of the message is not known to
	// this implementation.
	ErrUnknownKind Error = "unknown message kind"
	// ErrUnknownCommandKind indicates that the kind of command is not known
	// to this implemenatation.
	ErrUnknownCommandKind Error = "unknown command kind"
	// ErrNilCommand indicates that the command variable found is nil.
	ErrNilCommand Error = "nil command found"
)
