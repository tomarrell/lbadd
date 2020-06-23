package engine

import "fmt"

// Result represents an evaluation result of the engine. It is a mxn-matrix,
// where m and n is variable and depends on the passed in command.
type Result interface {
	Cols() []Column
	Rows() []Row
	fmt.Stringer
}

type IndexedGetter interface {
	Get(int) Value
}

type Sizer interface {
	Size() int
}

type Column interface {
	Type() Type
	IndexedGetter
	Sizer
}

type Row interface {
	IndexedGetter
	Sizer
}
