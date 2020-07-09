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
var _ Command = (*Delete)(nil)
var _ Command = (*DropTable)(nil)
var _ Command = (*DropIndex)(nil)
var _ Command = (*DropTrigger)(nil)
var _ Command = (*DropView)(nil)
var _ Command = (*Update)(nil)
var _ Command = (*Insert)(nil)
var _ Command = (*Join)(nil)
var _ Command = (*Limit)(nil)

// Command describes a structure that can be executed by the database executor.
// Instead of using bytecode, we use a hierarchical structure for the executor.
// This is mainly to increase readability.
type Command interface {
	fmt.Stringer
}

//go:generate stringer -type=JoinType

// JoinType is a type of join.
type JoinType uint8

// Known join types.
const (
	JoinUnknown JoinType = iota
	JoinLeft
	JoinLeftOuter
	JoinInner
	JoinCross
)

//go:generate stringer -type=UpdateOr

// UpdateOr is the type of update alternative that is specified in an update
// statement.
type UpdateOr uint8

// Known UpdateOrs
const (
	UpdateOrUnknown UpdateOr = iota
	UpdateOrRollback
	UpdateOrAbort
	UpdateOrReplace
	UpdateOrFail
	UpdateOrIgnore
)

//go:generate stringer -type=InsertOr

// InsertOr is the type of insert alternative that is specified in an insert
// statement.
type InsertOr uint8

// Known InsertOrs
const (
	InsertOrUnknown InsertOr = iota
	InsertOrReplace
	InsertOrRollback
	InsertOrAbort
	InsertOrFail
	InsertOrIgnore
)

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
		Command
		_list()
	}

	// Scan instructs the executor to use the contents of the nested table. It
	// is up to the executor whether he performs a full table scan or applies
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
		QualifiedName() string
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

	// Delete instructs the executor to delete all datasets from a table, that
	// match the filter.
	Delete struct {
		// Table is the table to delete datasets from.
		Table Table
		// Filter is an expression that a dataset has to match in order to be
		// deleted. This must not be nil. If all datasets from the table have to
		// be deleted, the filter will be a constant true expression.
		Filter Expr
	}

	// drop instructs the executor to drop the component that is specified by
	// the schema and name defined in this command.
	drop struct {
		// IfExists determines whether the executor should ignore an error that
		// occurs if the component with the defined name doesn't exist.
		IfExists bool
		// Schema is the schema of the referenced component.
		Schema string
		// Name is the name of the referenced component.
		Name string
	}

	// DropTable instructs the executor to drop the table with the name and
	// schema defined in this command.
	DropTable drop
	// DropView instructs the executor to drop the view with the name and schema
	// defined in this command.
	DropView drop
	// DropIndex instructs the executor to drop the index with the name and
	// schema defined in this command.
	DropIndex drop
	// DropTrigger instructs the executor to drop the trigger with the name and
	// schema defined in this command.
	DropTrigger drop

	// Update instructs the executor to update all datasets, for which the
	// filter expression evaluates to true, with the defined updates.
	Update struct {
		// UpdateOr is the OR clause in an update statement.
		UpdateOr UpdateOr
		// Table is the table on which the update should be performed.
		Table Table
		// Updates is a list of updates that must be applied.
		Updates []UpdateSetter
		// Filter is the filter expression, that determines, which datasets are
		// to be updated.
		Filter Expr
	}

	// UpdateSetter is an update that can be applied to a value in a dataset.
	UpdateSetter struct {
		// Cols are the columns of the dataset, that have to be updated. Because
		// the columns must be inside the table that this update refers to, and
		// no alias can be specified according to the grammar, this is just a
		// string, and not a full blown column.
		Cols []string
		// Value is the expression that the columns have to be set to.
		Value Expr
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
		// Natural indicates whether this join is a natural one.
		Natural bool
		// Type is the type of join that this join is.
		Type JoinType
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
		Limit Expr
		// Input is the input list of datasets.
		Input List
	}

	// Offset instructs to executor to skip the first Offset datasets from the
	// input list and return that truncated list. When used together with Limit,
	// please notice that the function composition (Limit ∘ Offset)(x) is not
	// commutative.
	Offset struct {
		// Offset is the amount of datasets that should be skipped from the
		// input list.
		Offset Expr
		// Input is the input list to truncate.
		Input List
	}

	// Empty instructs the executor to consider an empty list of datasets.
	Empty struct {
		// Cols are the columns in this empty list. This may be empty to
		// indicate a completely empty list.
		Cols []Column
	}

	// Distinct skips datasets from the list that already have been encountered
	// and returns a list with only distinct entries.
	Distinct struct {
		// Input is the input list that is filtered.
		Input List
	}

	// Values returns a list of datasets from the evaluated expressions.
	Values struct {
		// Values are the values that represent the datasets in this list. Each
		// dataset consists of all expressions that are in the dataset.
		Values [][]Expr
	}

	// Insert instructs the executor to insert the input list into the
	// specified table. The specified columns must match the columns from the
	// input list.
	Insert struct {
		// InsertOr is the specified fallback to perform when the insertion
		// fails.
		InsertOr InsertOr
		// Table is the specified table, where the input list is inserted into.
		Table Table
		// Cols are the columns, which are modified. The columns of the input
		// list have to match these columns.
		Cols []Column
		// DefaultValues determines whether to insert default values for all
		// (specified) columns. If this is set to true, the input list must not
		// be present.
		DefaultValues bool
		// Input is the input list of datasets, that will be inserted.
		Input List
	}
)

func (Scan) _list()     {}
func (Select) _list()   {}
func (Project) _list()  {}
func (Join) _list()     {}
func (Limit) _list()    {}
func (Offset) _list()   {}
func (Distinct) _list() {}
func (Values) _list()   {}

func (SimpleTable) _table() {}

// QualifiedName returns '<Schema>.<TableName>', or only '<TableName>' if no
// schema is specified.
func (t SimpleTable) QualifiedName() string {
	qualifiedName := t.Table
	if t.Schema != "" {
		qualifiedName = t.Schema + "." + qualifiedName
	}
	return qualifiedName
}

func (e Explain) String() string {
	return fmt.Sprintf("explanation: %v", e.Command)
}

func (s Scan) String() string {
	return fmt.Sprintf("Scan[table=%v]()", s.Table)
}

func (s Select) String() string {
	return fmt.Sprintf("Select[filter=%v](%v)", s.Filter, s.Input)
}

func (p Project) String() string {
	colStrs := make([]string, len(p.Cols))
	for i, col := range p.Cols {
		colStrs[i] = col.String()
	}
	return fmt.Sprintf("Project[cols=%v](%v)", strings.Join(colStrs, ","), p.Input)
}

func (d Delete) String() string {
	return fmt.Sprintf("Delete[filter=%v](%v)", d.Filter, d.Table)
}

func (d DropTable) String() string {
	table := d.Name
	if d.Schema != "" {
		table = d.Schema + "." + table
	}
	return fmt.Sprintf("DropTable[table=%v,ifexists=%v]()", table, d.IfExists)
}

func (d DropIndex) String() string {
	index := d.Name
	if d.Schema != "" {
		index = d.Schema + "." + index
	}
	return fmt.Sprintf("DropIndex[index=%v,ifexists=%v]()", index, d.IfExists)
}

func (d DropTrigger) String() string {
	trigger := d.Name
	if d.Schema != "" {
		trigger = d.Schema + "." + trigger
	}
	return fmt.Sprintf("DropTrigger[trigger=%v,ifexists=%v]()", trigger, d.IfExists)
}

func (d DropView) String() string {
	view := d.Name
	if d.Schema != "" {
		view = d.Schema + "." + view
	}
	return fmt.Sprintf("DropView[view=%v,ifexists=%v]()", view, d.IfExists)
}

func (u Update) String() string {
	var sets []string
	for _, set := range u.Updates {
		sets = append(sets, set.String())
	}
	return fmt.Sprintf("Update[or=%v,table=%v,sets=(%v),filter=%v]", u.UpdateOr, u.Table, strings.Join(sets, ","), u.Filter)
}

func (u UpdateSetter) String() string {
	return fmt.Sprintf("(%v)=%v", strings.Join(u.Cols, ","), u.Value)
}

func (c Column) String() string {
	if c.Alias == "" {
		return c.Column.String()
	}
	return fmt.Sprintf("%v AS %v", c.Column, c.Alias)
}

func (j Join) String() string {
	var buf strings.Builder
	// configuration
	var cfg []string
	if j.Filter != nil {
		cfg = append(cfg, fmt.Sprintf("filter=%v", j.Filter))
	}
	if j.Natural {
		cfg = append(cfg, fmt.Sprintf("natural=%v", j.Natural))
	}
	if j.Type != JoinUnknown {
		cfg = append(cfg, fmt.Sprintf("type=%v", j.Type))
	}
	// compose
	buf.WriteString(fmt.Sprintf("Join[%s](%v,%v)", strings.Join(cfg, ","), j.Left, j.Right))
	return buf.String()
}

func (l Limit) String() string {
	return fmt.Sprintf("Limit[limit=%v](%v)", l.Limit, l.Input)
}

func (o Offset) String() string {
	return fmt.Sprintf("Offset[offset=%v](%v)", o.Offset, o.Input)
}

func (e Empty) String() string {
	colStrs := make([]string, len(e.Cols))
	for i, col := range e.Cols {
		colStrs[i] = col.String()
	}
	return fmt.Sprintf("Empty[cols=%v]()", strings.Join(colStrs, ","))
}

func (d Distinct) String() string {
	return fmt.Sprintf("Distinct[](%v)", d.Input.String())
}

func (t SimpleTable) String() string {
	var buf strings.Builder
	if t.Schema != "" {
		buf.WriteString(t.Schema + ".")
	}
	buf.WriteString(t.Table)
	if t.Alias != "" {
		buf.WriteString(" AS " + t.Alias)
	}
	return buf.String()
}

func (v Values) String() string {
	var values []string
	for _, val := range v.Values {
		var exprs []string
		for _, expr := range val {
			exprs = append(exprs, expr.String())
		}
		values = append(values, "("+strings.Join(exprs, ",")+")")
	}
	return fmt.Sprintf("Values[](%v)", strings.Join(values, ","))
}

func (i Insert) String() string {
	var cols []string
	for _, col := range i.Cols {
		cols = append(cols, col.String())
	}
	return fmt.Sprintf("Insert[table=%v,cols=%v](%v)", i.Table, strings.Join(cols, ","), i.Input)
}
