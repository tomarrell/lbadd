package command

// Expr is a structure that represents an SQL expression.
type Expr interface{}

type (
	// LiteralExpr is a representation of an expr with the production
	//
	//	expr : literal-value
	LiteralExpr struct {
		// LiteralValue is the string value that was extracted from the parser. It
		// is a plain string and was not modified, because the executor must
		// determine which data type it must be interpreted at.
		LiteralValue string
	}

	// EqualityExpr is a representation of an expr with the production
	//
	//	expr : expr IS NOT? expr
	EqualityExpr struct {
		// Left is the left hand expression of this equality expression.
		Left Expr
		// Right is the right hand expression of this equality expression.
		Right Expr
		// Invert is a flag indicating that this equality expression must be
		// interpreted as un-equality expression.
		Invert bool
	}

	// RangeExpr is a representation of an expr with the production
	//
	//	expr : expr IS NOT? BETWEEN expr AND expr
	RangeExpr struct {
		// Lo  is the lower bound of this range.
		Lo Expr
		// Hi is the upper bound of this range.
		Hi Expr
		// Invert is a flag indicating that this range expression indicates a
		// range that a value must NOT match.
		Invert bool
	}
)
