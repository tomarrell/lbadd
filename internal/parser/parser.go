package parser

import (
	"github.com/tomarrell/lbadd/internal/parser/ast"
)

// Error allows constant errors.
type Error string

func (s Error) Error() string { return string(s) }

// parser errors
const (
	ErrIncompleteStatement  = Error("incomplete statement")
	ErrPrematureEOF         = Error("unexpectedly reached EOF")
	ErrScanner              = Error("scanner")
	ErrUnexpectedToken      = Error("unexpected token")
	ErrUnknownToken         = Error("unknown token")
	ErrUnsupportedConstruct = Error("unsupported construct")
)

// Parser describes a parser that returns (maybe multiple) SQLStatements from a
// given input.
type Parser interface {
	// Next returns stmt=<statement>, errs=nil, ok=true if a statement was
	// parsed successfully without any parse errors. If there were parse errors,
	// Next will return stmt=<statement>, errs=([]error), ok=true.
	//
	// stmt always is the statement that was parsed. If it could not be parsed
	// free of errors, the statement might be incomplete or incorrect, but
	// efforts will be taken to parse as much out of the given input as
	// possible. ok indicates whether any statement could have been parsed, or
	// more precisely, if the underlying scanner had any more tokens.
	//
	// If ok=false, that means that the parser has reached its EOF and no more
	// statements can be returned. Subsequent calls to Next will result in
	// stmt=nil, errs=nil, ok=false.
	Next() (stmt *ast.SQLStmt, errs []error, ok bool)
}

// New creates a new, ready to use parser, that will parse the given input
// string.
func New(input string) Parser {
	return NewSimpleParser(input)
}
