package scanner

import (
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

// Error allows constant errors.
type Error string

func (s Error) Error() string { return string(s) }

// Constant errors
const (
	ErrUnexpectedToken = Error("unexpected token")
)

// Scanner is the interface that describes a scanner.
type Scanner interface {
	Next() token.Token
	Peek() token.Token
}
