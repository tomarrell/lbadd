package command

import "fmt"

type (
	// Expr is a marker interface for anything that is an expression. Different
	// implementations of this interface represent different productions of the
	// expression rule in the SQL grammar.
	Expr interface {
		fmt.Stringer
		_expr()
	}

	// LiteralExpr is a simple literal expression that has a single string
	// value.
	LiteralExpr struct {
		// Value is the simple string value of this expression.
		Value string
	}

	// EqualityExpr is an expression with a left and right side expression, and
	// represents the condition that both expressions are equal. If this
	// equality expression is inverted, the condition is, that both sides are
	// un-equal.
	EqualityExpr struct {
		// Left is the left hand side expression.
		Left Expr
		// Right is the right hand side expression.
		Right Expr
		// Invert determines whether this equality expression must be considered
		// as in-equality expression.
		Invert bool
	}

	// RangeExpr is an expression with a needle, an upper and a lower bound. It
	// must be evaluated to true, if needle is within the lower and upper bound,
	// or if the needle is not between the bounds and the range is inverted.
	RangeExpr struct {
		// Needle is the value that is evaluated if it is between Lo and Hi.
		Needle Expr
		// Lo is the lower bound of this range.
		Lo Expr
		// Hi is the upper bound of this range.
		Hi Expr
		// Invert determines if Needle must be between or not between the bounds
		// of this range.
		Invert bool
	}
)

func (LiteralExpr) _expr()  {}
func (EqualityExpr) _expr() {}
func (RangeExpr) _expr()    {}

func (l LiteralExpr) String() string {
	return l.Value
}

func (e EqualityExpr) String() string {
	if e.Invert {
		return fmt.Sprintf("%v!=%v", e.Left, e.Right)
	}
	return fmt.Sprintf("%v==%v", e.Left, e.Right)
}

func (r RangeExpr) String() string {
	if r.Invert {
		return fmt.Sprintf("![%v;%v]", r.Lo, r.Hi)
	}
	return fmt.Sprintf("[%v;%v]", r.Lo, r.Hi)
}
