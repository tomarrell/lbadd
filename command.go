package lbadd

type command int

const (
	commandUnknown command = iota
	commandInsert
	commandSelect
	commandDelete
	commandCreateTable
)

func newCommand(cmd string) command {
	switch toUp(cmd) {
	case commandInsert.String():
		return commandInsert
	case commandSelect.String():
		return commandSelect
	case commandDelete.String():
		return commandDelete
	case commandCreateTable.String():
		return commandCreateTable
	default:
		return commandUnknown
	}
}

func (c command) String() string {
	switch c {
	case commandInsert:
		return "INSERT"
	case commandSelect:
		return "SELECT"
	case commandDelete:
		return "DELETE"
	case commandCreateTable:
		return "CREATE TABLE"
	default:
		return "UNKNOWN"
	}
}
