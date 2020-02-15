package token

import "fmt"

// Token describes a single token that was produced by a scanner.
type Token interface {
	Positioner
	Lengther
	Typer
	Valuer
}

func Equal(t1, t2 Token) bool {
	return t1.Line() == t2.Line() &&
		t1.Col() == t2.Col() &&
		t1.Offset() == t2.Offset() &&
		t1.Length() == t2.Length() &&
		t1.Type() == t2.Type() &&
		t1.Value() == t2.Value()
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
