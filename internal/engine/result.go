package engine

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/engine/types"
)

// Result represents an evaluation result of the engine. It is a mxn-matrix,
// where m and n is variable and depends on the passed in command.
type Result interface {
	Cols() []Column
	Rows() []Row
	fmt.Stringer
}

// IndexedGetter wraps a Get(index) method.
type IndexedGetter interface {
	Get(int) types.Value
}

// Sizer wraps the basic Size() method.
type Sizer interface {
	Size() int
}

// Column is an iterator over cells in a column. All of the cells must have the
// same type.
type Column interface {
	Type() types.Type
	IndexedGetter
	Sizer
}

// Row is an iterator over cells in a row. The cells may have different types.
type Row interface {
	IndexedGetter
	Sizer
}
