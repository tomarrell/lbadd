package command

var _ Command = (*Select)(nil)

// Select describes a command that represents a select statement.
type Select struct {
	// Tables are the tables that this select command must read from.
	Tables []string
	// Cols are the columns that go into the result table.
	Cols  []string
	Where Expr
}

// Type returns TypeSelect.
func (s Select) Type() Type { return TypeSelect }
