package message

import "github.com/tomarrell/lbadd/internal/compiler/command"

var _ Message = (*Command)(nil)
var _ command.Command = (*Command)(nil)

// NewCommand creates a new Command variable.
func NewCommand() *Command {
	return &Command{}
}

// Kind returns KindCommand.
func (*Command) Kind() Kind {
	return KindCommand
}

// Kind returns KindExpr
func (*Expr) Kind() Kind {
	return KindExpr
}

// Kind returns KindCommandScan.
func (*Command_Scan) Kind() Kind {
	return KindCommandScan
}

// Kind returns KindCommandSelect.
func (*Command_Select) Kind() Kind {
	return KindCommandSelect
}

// Kind returns KindCommandProject.
func (*Command_Project) Kind() Kind {
	return KindCommandProject
}

// Kind returns KindCommandDelete.
func (*Command_Delete) Kind() Kind {
	return KindCommandDelete
}

// Kind returns KindCommandUpdate.
func (*Command_Update) Kind() Kind {
	return KindCommandUpdate
}

// Kind returns KindCommandDrop.
func (*CommandDrop) Kind() Kind {
	return KindCommandDrop
}

// Kind returns KindCommandLimit.
func (*Command_Limit) Kind() Kind {
	return KindCommandLimit
}

// Kind returns KindCommandJoin.
func (*Command_Join) Kind() Kind {
	return KindCommandJoin
}

// Kind returns KindCommandInsert.
func (*Command_Insert) Kind() Kind {
	return KindCommandInsert
}
