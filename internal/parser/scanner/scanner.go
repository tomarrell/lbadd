package scanner

import (
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

// sentinel allows constant errors
type sentinel string

func (s sentinel) Error() string { return string(s) }

// Constant errors
const (
	ErrUnexpectedToken = sentinel("unexpected token")
)

// Scanner is the interface that describes a scanner.
type Scanner interface {
	Next() token.Token
	Peek() token.Token
}
