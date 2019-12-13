package main

type command int

const (
	commandUnknown command = iota
	commandInsert
	commandSelect
	commandDelete
)

func newCommand(cmd string) command {
	switch cmd {
	case commandInsert.String():
		return commandInsert
	case commandSelect.String():
		return commandInsert
	case commandDelete.String():
		return commandDelete
	default:
		return commandUnknown
	}
}

func (c command) String() string {
	switch c {
	case commandInsert:
		return "insert"
	case commandSelect:
		return "select"
	case commandDelete:
		return "delete"
	default:
		return "unknown"
	}
}
