package driver

type Error string

func (e Error) Error() string { return string(e) }

// Constant errors
const (
	ErrConnectionClosed = Error("connection is closed")
)
