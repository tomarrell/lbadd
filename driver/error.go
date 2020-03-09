package driver

// Error provides constant errors to the driver package.
type Error string

func (e Error) Error() string { return string(e) }

// Constant errors
const (
	ErrConnectionClosed = Error("connection is closed")
)
