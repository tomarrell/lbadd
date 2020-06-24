package result

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/tomarrell/lbadd/internal/engine/types"
)

type valueResult struct {
	vals [][]types.Value
}

type valueColumn struct {
	valueResult
	colIndex int
}

type valueRow struct {
	valueResult
	rowIndex int
}

// FromValues wraps the given values in a result, performing a type check on all
// columns. If any column does not have a consistent type, an error will be
// returned, together with NO result.
func FromValues(vals [][]types.Value) (Result, error) {
	for x := 0; x < len(vals[0]); x++ { // cols
		t := vals[0][x].Type()
		for y := 0; y < len(vals); y++ { // rows
			if !vals[y][x].Is(t) {
				return nil, types.ErrTypeMismatch(t, vals[y][x].Type())
			}
		}
	}
	return valueResult{
		vals: vals,
	}, nil
}

func (r valueResult) Cols() []Column {
	result := make([]Column, 0)
	for i := 0; i < len(r.vals[0]); i++ {
		result = append(result, valueColumn{
			valueResult: r,
			colIndex:    i,
		})
	}
	return result
}

func (r valueResult) Rows() []Row {
	result := make([]Row, 0)
	for i := 0; i < len(r.vals); i++ {
		result = append(result, valueRow{
			valueResult: r,
			rowIndex:    i,
		})
	}
	return result
}

func (r valueResult) String() string {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 1, 3, ' ', 0)

	var types []string
	for _, col := range r.Cols() {
		types = append(types, col.Type().String())
	}
	_, _ = fmt.Fprintln(w, strings.Join(types, "\t"))

	for _, row := range r.Rows() {
		var strVals []string
		for i := 0; i < row.Size(); i++ {
			strVals = append(strVals, row.Get(i).String())
		}
		_, _ = fmt.Fprintln(w, strings.Join(strVals, "\t"))
	}
	_ = w.Flush()
	return buf.String()
}

func (c valueColumn) Type() types.Type {
	return c.Get(0).Type()
}

func (c valueColumn) Get(index int) types.Value {
	return c.valueResult.vals[index][c.colIndex]
}

func (c valueColumn) Size() int {
	return len(c.valueResult.vals)
}

func (r valueRow) Get(index int) types.Value {
	return r.valueResult.vals[r.rowIndex][index]
}

func (r valueRow) Size() int {
	return len(r.valueResult.vals[0])
}
