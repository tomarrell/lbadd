package cmp

import (
	"strings"

	"github.com/tomarrell/lbadd/internal/parser/ast"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

//go:generate stringer -type=DeltaType

// DeltaType describes the type of a delta, for example Nilness or TokenValue.
type DeltaType uint16

const (
	// Unknown means, that the type in the delta has not been set properly. This
	// must never be used. If you encounter this as a value, a developer has
	// made an error. Please open an issue in that case.
	Unknown DeltaType = iota
	// Nilness describes, that two components in the same position (in the AST)
	// have different nilness, i.e. one component is nil, the other is not.
	Nilness
	// TokenValue describes, that two tokens in the same position (in the AST)
	// have different values.
	TokenValue
	// TokenPosition describes, that two tokens in the same position (in the
	// AST) have different location information.
	TokenPosition
)

type Delta struct {
	// Path indicates the path within the struct, where the delta is present,
	// for example:
	//
	//  SQLStmt.AlterTableStmt.ColumnDef.ColumnName
	//
	// This would indicate, that the delta describes a difference in the column
	// name of the alter table statement of the sql statement. Effectively, this
	// means, that the tokens in ColumnName in the first and second SQLStmt
	// differ. See the message and the type (Typ) for additional difference
	// information.
	Path string
	// Typ describes the type of this delta. For information on what type means
	// what exactly, please refer to the documentation of the respective type.
	Typ DeltaType
	// Message contains a detailed description of what the delta describes, for
	// example, if the difference is in a token value, or nilness of a member.
	// This message is intended to be human readable, not machine processable.
	Message string
	// Left is the left hand side component of this delta.
	Left interface{}
	// Right is the right hand side component of this delta.
	Right interface{}
}

// CompareAST compares two ASTs (specifically, two (*ast.SQLStmt)s) against each
// other, and returns a list of deltas, which will be nil if the ASTs are equal.
func CompareAST(left, right *ast.SQLStmt) (deltas []Delta) {
	return compareSQLStmt(left, right, append(path{}, "SQLStmt"))
}

type path []string

func (p path) String() string { return strings.Join(p, ".") }

func compareSQLStmt(left, right *ast.SQLStmt, path path) (deltas []Delta) {
	deltas = append(deltas, compareTokens(left.Explain, right.Explain, append(path, "Explain"))...)
	deltas = append(deltas, compareTokens(left.Query, right.Query, append(path, "Query"))...)
	deltas = append(deltas, compareTokens(left.Plan, right.Plan, append(path, "Plan"))...)
	deltas = append(deltas, compareAlterTableStmt(left.AlterTableStmt, right.AlterTableStmt, append(path, "AlterTableStmt"))...)
	// TODO(TimSatke) all other fields
	return
}

func compareAlterTableStmt(left, right *ast.AlterTableStmt, path path) (deltas []Delta) {
	deltas = append(deltas, compareTokens(left.Alter, right.Alter, append(path, "Alter"))...)
	deltas = append(deltas, compareTokens(left.Table, right.Table, append(path, "Table"))...)
	deltas = append(deltas, compareTokens(left.SchemaName, right.SchemaName, append(path, "SchemaName"))...)
	deltas = append(deltas, compareTokens(left.Period, right.Period, append(path, "Period"))...)
	deltas = append(deltas, compareTokens(left.TableName, right.TableName, append(path, "TableName"))...)
	deltas = append(deltas, compareTokens(left.Rename, right.Rename, append(path, "Rename"))...)
	deltas = append(deltas, compareTokens(left.To, right.To, append(path, "To"))...)
	deltas = append(deltas, compareTokens(left.NewTableName, right.NewTableName, append(path, "NewTableName"))...)
	deltas = append(deltas, compareTokens(left.Column, right.Column, append(path, "Column"))...)
	deltas = append(deltas, compareTokens(left.ColumnName, right.ColumnName, append(path, "ColumnName"))...)
	deltas = append(deltas, compareTokens(left.NewColumnName, right.NewColumnName, append(path, "NewColumnName"))...)
	deltas = append(deltas, compareTokens(left.Add, right.Add, append(path, "Add"))...)
	// TODO(TimSatke) compare column def
	return
}

func compareTokens(left, right token.Token, path path) (deltas []Delta) {
	if (left == nil && right != nil) ||
		(left != nil && right == nil) {
		deltas = append(deltas, Delta{
			Path:    path.String(),
			Typ:     Nilness,
			Message: "one token was nil while the other one wasn't",
			Left:    left,
			Right:   right,
		})
	}

	if left == right {
		return
	}

	if left.Line() != right.Line() {
		deltas = append(deltas, Delta{
			Path:    path.String(),
			Typ:     TokenPosition,
			Message: "lines don't match",
			Left:    left.Line(),
			Right:   right.Line(),
		})
	}

	if left.Col() != right.Col() {
		deltas = append(deltas, Delta{
			Path:    path.String(),
			Typ:     TokenPosition,
			Message: "cols don't match",
			Left:    left.Col(),
			Right:   right.Col(),
		})
	}

	if left.Offset() != right.Offset() {
		deltas = append(deltas, Delta{
			Path:    path.String(),
			Typ:     TokenPosition,
			Message: "offsets don't match",
			Left:    left.Offset(),
			Right:   right.Offset(),
		})
	}

	if left.Length() != right.Length() {
		deltas = append(deltas, Delta{
			Path:    path.String(),
			Typ:     TokenValue,
			Message: "lengths don't match",
			Left:    left.Length(),
			Right:   right.Length(),
		})
	}

	if left.Type() != right.Type() {
		deltas = append(deltas, Delta{
			Path:    path.String(),
			Typ:     TokenValue,
			Message: "types don't match",
			Left:    left.Type(),
			Right:   right.Type(),
		})
	}

	if left.Value() != right.Value() {
		deltas = append(deltas, Delta{
			Path:    path.String(),
			Typ:     TokenValue,
			Message: "values don't match",
			Left:    left.Value(),
			Right:   right.Value(),
		})
	}

	return
}
