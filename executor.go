package lbadd

import "fmt"

// Contains a command and associated information required to execute such command
type instruction struct {
	command command
	table   string
	params  []string
}

// A single column on a single row
type record []byte

// A row containing multiple records
type row []record

// A response from the executor
type result struct {
	columns      []column // columns in order of stored rows
	rows         []row    // the set of rows
	rowsAffected int      // the number of rows affected by execution
	created      int      // the number of resources created
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

// The executor takes an instruction, and coordinates the operations which are
// required to fulfill the instruction, executing these against the DB. It
// also returns the result of the instruction.
func (e *executor) execute(instr instruction) (result, error) {
	switch instr.command {
	case commandInsert:
		return result{}, fmt.Errorf("unimplmented")
	case commandSelect:
		return result{}, fmt.Errorf("unimplmented")
	case commandDelete:
		return result{}, fmt.Errorf("unimplmented")
	case commandCreateTable:
		e.db.tables[instr.table] = table{}
		return result{}, fmt.Errorf("unimplmented")
	default:
		return result{}, fmt.Errorf("invalid executor command")
	}
}
