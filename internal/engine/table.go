package engine

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/tomarrell/lbadd/internal/engine/types"
)

var (
	// EmptyTable is the empty table, with 0 cols and 0 rows.
	EmptyTable = Table{
		Cols: make([]Col, 0),
		Rows: make([]Row, 0),
	}
)

// Table is a one-dimensional collection of Rows.
type Table struct {
	Cols []Col
	Rows []Row
}

// Col is a header for a single column in a table, containing the qualified name
// of the col, a possible alias and the col data type.
type Col struct {
	QualifiedName string
	Alias         string
	Type          types.Type
}

// Row is a one-dimensional collection of values.
type Row struct {
	Values []types.Value
}

// RemoveColumnByQualifiedName will remove the first column with the given
// qualified name from the table, and return the new table. The original table
// will not be modified. If no such column exists, the original table is
// returned.
func (t Table) RemoveColumnByQualifiedName(qualifiedName string) Table {
	index := -1
	for i, col := range t.Cols {
		if col.QualifiedName == qualifiedName {
			index = i
			break
		}
	}
	if index != -1 {
		return t.RemoveColumn(index)
	}
	return t
}

// HasColumn inspects the table's columns and determines whether the table has
// any column, that has the given name as qualified name OR as alias.
func (t Table) HasColumn(qualifiedNameOrAlias string) bool {
	for _, col := range t.Cols {
		if col.QualifiedName == qualifiedNameOrAlias || col.Alias == qualifiedNameOrAlias {
			return true
		}
	}
	return false
}

// RemoveColumn works on a copy of the table, and removes the column with the
// given index from the copy. After removal, the copy is returned.
func (t Table) RemoveColumn(index int) Table {
	t.Cols = append(t.Cols[:index], t.Cols[index+1:]...)
	for i := range t.Rows {
		t.Rows[i].Values = append(t.Rows[i].Values[:index], t.Rows[i].Values[index+1:]...)
	}
	return t
}

func (t Table) String() string {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 1, 3, ' ', 0)

	var colNames []string
	for _, col := range t.Cols {
		colName := col.QualifiedName
		if col.Alias != "" {
			colName = col.Alias
		}
		colNames = append(colNames, colName+" ("+col.Type.String()+")")
	}
	_, _ = fmt.Fprintln(w, strings.Join(colNames, "\t"))

	for _, row := range t.Rows {
		var strVals []string
		for i := 0; i < len(row.Values); i++ {
			strVals = append(strVals, row.Values[i].String())
		}
		_, _ = fmt.Fprintln(w, strings.Join(strVals, "\t"))
	}
	_ = w.Flush()
	return buf.String()
}
