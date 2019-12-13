package main

import "fmt"

// Contains a command and associated information required to execute such command
type instruction struct {
	command command
}

// Execute executes an instruction against the database
type executor struct {
	db *db
}

func newExecutor() *executor {
	return &executor{
		db: newDB(),
	}
}

func (e *executor) execute(instr instruction) error {
	switch instr.command {
	case commandInsert:
	case commandSelect:
	case commandDelete:
	default:
		return fmt.Errorf("invalid executor command")
	}

	return nil
}
