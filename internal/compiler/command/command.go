package command

// The implemented command structure is inspired by the QIR proposed by the
// following paper. https://arxiv.org/pdf/1607.04197.pdf

import (
	"fmt"
	"strings"
)

var _ Command = (*Explain)(nil)
var _ Command = (*Scan)(nil)
var _ Command = (*Select)(nil)
var _ Command = (*Project)(nil)
var _ Command = (*Join)(nil)
var _ Command = (*Limit)(nil)

// Command describes a structure that can be executed by the database executor.
// Instead of using bytecode, we use a hierarchical structure for the executor.
// This is mainly to increase readability.
type Command interface {
	fmt.Stringer
}

type (
	// Explain instructs the executor to explain the nested command instead of
	// executing it.
	Explain struct {
		// Command is the command that will be explained, but not executed.
		Command Command
	}

	// List is a marker interface that facilitates creating a type hierarchy for
	// the command model.
	List interface {
		fmt.Stringer
		_list()
	}

	// Scan instructs the executor to use the contents of the nested table. If
	// it up to the executor whether he performs a full table scan or applies
	// possible optimizations, like a search through indices.
	Scan struct {
		// Table is the table whose contents will be used.
		Table Table
	}

	// Table is a marker interface, allowing for different specifications of
	// tables, such as a simple table, specified by schema and table name, or a
	// more sophisticated table, such as a combination of multiple sub-tables or
	// select statements.
	Table interface {
		_table()
	}

	// SimpleTable is a table that is only specified by schema and table name,
	// and an optional alias. It is also optionally indexed by an index.
	//
	// SimpleTable represents the first grammar production of table-or-subquery.
	SimpleTable struct {
		// Schema name of the table. May be empty.
		Schema string
		// Table name of this table. Since this is a simple table, the table
		// name is a string and not an expression. Use other Table
		// implementations for more complex tables.
		Table string
		// Alias name of this table. May be empty.
		Alias string
		// Indexed indicates, whether this table is indexed by an index. If this
		// is false, Index must be the empty string.
		Indexed bool
		// Index is the name of the index that indexed this table, or empty, if
		// Indexed is false.
		Index string
	}

	// Select represents a selection that should be performed by the executor
	// over the nested input. Additionally, a filter can be specified which must
	// be respected by the executor.
	Select struct {
		// Filter is an expression that filters elements in this selection to
		// only elements, that pass this filter.
		Filter Expr
		// Input is the input list over which the selection takes place.
		Input List
	}

	// Project represents a projection that should be performed by the executor
	// over the nested input. The projected columns are specified in
	// (command.Project).Cols.
	Project struct {
		// Cols are the columns that this projection projects. Most of the time,
		// this will be the columns from the SELECT statement.
		Cols []Column
		// Input is the input list over which the projection takes place.
		Input List
	}

	// Column represents a database table column.
	Column struct {
		// Table is the table name that this column belongs to. May be empty, as
		// this is a representation derived from the AST. If this is empty, the
		// executor has to interpolate the table from the execution context.
		Table string
		// Column is the name of the column.
		Column Expr
		// Alias is the alias name for this table. May be empty.
		Alias string
	}

	// Join instructs the executor to produce a list from the left and right
	// input list. Lists are merged with respect to the given filter.
	Join struct {
		// Filter defines the condition that has to apply to two datasets from
		// the left and right list in order to be merged.
		Filter Expr
		// Left is the left input list.
		Left List
		// Right is the right input list.
		Right List
	}

	// Limit instructs the executor to only respect the first Limit datasets
	// from the input list.
	Limit struct {
		// Limit is the amount of datasets that are respected from the input
		// list (top to bottom).
		Limit uint64
		// Input is the input list of datasets.
		Input List
	}
)

func (Scan) _list()    {}
func (Select) _list()  {}
func (Project) _list() {}
func (Join) _list()    {}
func (Limit) _list()   {}

func (SimpleTable) _table() {}

func (e Explain) String() string {
	return fmt.Sprintf("explanation: %v", e.Command)
}

func (s Scan) String() string {
	return fmt.Sprintf("Scan[table=%v]()", s.Table)
}

func (s Select) String() string {
	if s.Filter == nil {
		return fmt.Sprintf("Select[](%v)", s.Input)
	}
	return fmt.Sprintf("Select[filter=%v](%v)", s.Filter, s.Input)
}

func (p Project) String() string {
	colStrs := make([]string, len(p.Cols))
	for i, col := range p.Cols {
		colStrs[i] = col.String()
	}
	return fmt.Sprintf("Project[cols=%v](%v)", strings.Join(colStrs, ","), p.Input)
}

func (c Column) String() string {
	if c.Alias == "" {
		return c.Column.String()
	}
	return fmt.Sprintf("%v AS %v", c.Column, c.Alias)
}

func (j Join) String() string {
	if j.Filter == nil {
		return fmt.Sprintf("Join[](%v,%v)", j.Left, j.Right)
	}
	return fmt.Sprintf("Join[filter=%v](%v,%v)", j.Filter, j.Left, j.Right)
}

func (l Limit) String() string {
	return fmt.Sprintf("Limit[limit=%d](%v)", l.Limit, l.Input)
}
