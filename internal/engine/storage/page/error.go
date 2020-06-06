package page

// Error is a sentinel error.
type Error string

func (e Error) Error() string { return string(e) }

// Sentinel errors.
const (
	ErrUnknownHeader   = Error("unknown header")
	ErrInvalidPageSize = Error("invalid page size")
	ErrPageTooSmall    = Error("page is too small for the requested data")
)
