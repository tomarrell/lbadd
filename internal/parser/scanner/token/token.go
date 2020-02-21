package token

import "fmt"

// Token describes a single token that was produced by a scanner.
type Token interface {
	Positioner
	Lengther
	Typer
	Valuer
}

// ErrorToken describes a single token that was produced by the scanner and
// holds an error object.
type ErrorToken interface {
	Token
	error
}

// Positioner describes something that has a 1-based line and col, and a 0-based
// offset.
type Positioner interface {
	Line() int
	Col() int
	Offset() int
}

// Lengther describes something that has a length.
type Lengther interface {
	Length() int
}

// Typer describes a token that has a token type.
type Typer interface {
	Type() Type
}

// Valuer describes something that has a string value.
type Valuer interface {
	Value() string
}

var _ Token = (*tok)(nil) // ensure that tok implements Token

type tok struct {
	line, col int
	offset    int
	length    int
	typ       Type
	value     string
}

// New creates a new Token implementation, representing the given values.
func New(line, col, offset, length int, typ Type, value string) Token {
	return tok{
		line:   line,
		col:    col,
		offset: offset,
		length: length,
		typ:    typ,
		value:  value,
	}
}

func (t tok) Line() int {
	return t.line
}

func (t tok) Col() int {
	return t.col
}

func (t tok) Offset() int {
	return t.offset
}

func (t tok) Length() int {
	return t.length
}

func (t tok) Type() Type {
	return t.typ
}

func (t tok) Value() string {
	return t.value
}

func (t tok) String() string {
	return fmt.Sprintf("%s(%s)", t.typ.String(), t.value)
}

type errorTok struct {
	Token
	err error
}

// NewErrorToken creates a new token with the given attributes and the given
// error.
func NewErrorToken(line, col, offset, length int, typ Type, err error) Token {
	return errorTok{
		Token: New(line, col, offset, length, typ, err.Error()),
		err:   err,
	}
}

func (t errorTok) Error() string {
	return t.String()
}

func (t errorTok) String() string {
	return fmt.Sprintf("%s(%s)", t.Type().String(), t.err)
}
