package cmp

import (
	"reflect"
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
	return compare(left, right, path{})
}

type path []string

func (p path) String() string { return strings.Join(p, ".") }

func compare(left, right interface{}, parent path) (deltas []Delta) {
	leftVal, rightVal := reflect.ValueOf(left), reflect.ValueOf(right)
	leftElem := leftVal.Elem()
	rightElem := rightVal.Elem()

	if (leftVal.IsNil() && !rightVal.IsNil()) ||
		(!leftVal.IsNil() && rightVal.IsNil()) {
		// one of the values is nil
		which := "left"
		whichNot := "right"
		p := parent
		if rightElem.Interface() == nil {
			which = "right"
			whichNot = "left"
			p = append(p, leftElem.Type().Name())
		} else {
			p = append(p, rightElem.Type().Name())
		}

		deltas = append(deltas, Delta{
			Left:    leftElem.Interface,
			Right:   rightElem.Interface(),
			Path:    p.String(),
			Typ:     Nilness,
			Message: which + " was nil, while " + whichNot + " wasn't",
		})
	} else if leftVal.IsNil() && rightVal.IsNil() {
		return
	}

	// both incoming are not nil

	typ := leftElem.Type()
	if typ != rightElem.Type() {
		panic("struct types are not equal, thus not comparable")
	}

	path := append(parent, typ.Name())

	for i := 0; i < typ.NumField(); i++ {
		leftVal := reflect.ValueOf(left).Elem().Field(i).Interface()
		rightVal := reflect.ValueOf(right).Elem().Field(i).Interface()
		if (leftVal == nil && rightVal != nil) ||
			(leftVal != nil && rightVal == nil) {
			// only one is nil
			which := "left"
			whichNot := "right"
			if rightVal == nil {
				which = "right"
				whichNot = "left"
			}

			deltas = append(deltas, Delta{
				Left:    leftVal,
				Right:   rightVal,
				Path:    append(path, typ.Field(i).Name).String(),
				Typ:     Nilness,
				Message: which + " was nil, while " + whichNot + " wasn't",
			})
		} else if leftVal == nil && rightVal == nil {
			// both are nil, no-op
		} else {
			// both are not nil and we have to compare the values
			tok1, ok1 := leftVal.(token.Token)
			tok2, ok2 := rightVal.(token.Token)
			if ok1 && ok2 {
				deltas = append(deltas, compareToken(tok1, tok2, append(path, typ.Field(i).Name))...)
			} else {
				deltas = append(deltas, compare(leftVal, rightVal, path)...)
			}
		}
	}

	return
}

func compareToken(left, right token.Token, path path) (deltas []Delta) {
	if left.Col() != right.Col() {
		deltas = append(deltas, Delta{
			Left:    left,
			Right:   right,
			Path:    path.String(),
			Typ:     TokenPosition,
			Message: "difference in attribute 'Col'",
		})
	}
	if left.Length() != right.Length() {
		deltas = append(deltas, Delta{
			Left:    left,
			Right:   right,
			Path:    path.String(),
			Typ:     TokenPosition,
			Message: "difference in attribute 'Length'",
		})
	}
	if left.Line() != right.Line() {
		deltas = append(deltas, Delta{
			Left:    left,
			Right:   right,
			Path:    path.String(),
			Typ:     TokenPosition,
			Message: "difference in attribute 'Line'",
		})
	}
	if left.Offset() != right.Offset() {
		deltas = append(deltas, Delta{
			Left:    left,
			Right:   right,
			Path:    path.String(),
			Typ:     TokenPosition,
			Message: "difference in attribute 'Offset'",
		})
	}
	if left.Type() != right.Type() {
		deltas = append(deltas, Delta{
			Left:    left,
			Right:   right,
			Path:    path.String(),
			Typ:     TokenPosition,
			Message: "difference in attribute 'Type'",
		})
	}
	if left.Value() != right.Value() {
		deltas = append(deltas, Delta{
			Left:    left,
			Right:   right,
			Path:    path.String(),
			Typ:     TokenValue,
			Message: "difference in attribute 'Value'",
		})
	}
	return
}
