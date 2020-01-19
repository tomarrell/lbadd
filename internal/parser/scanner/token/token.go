package token

import "fmt"

type Token interface {
	Positioner
	Lengther
	Typer
	Valuer
}

type Positioner interface {
	Line() int
	Col() int
	Offset() int
}

type Lengther interface {
	Length() int
}

type Typer interface {
	Type() Type
}

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
