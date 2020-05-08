package message

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrUnknownKind Error = "unknown message kind"
)
