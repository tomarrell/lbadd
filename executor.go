package lbadd

import (
	"fmt"
	"regexp"

	"github.com/tomarrell/lbadd/internal/btree"
)

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

type exeConfig struct {
	order int
}

// Execute executes an instruction against the database
type executor struct {
	db  *db
	cfg exeConfig
}

func newExecutor(cfg exeConfig) *executor {
	return &executor{
		db:  newDB(),
		cfg: cfg,
	}
}

// The executor takes an instruction, and coordinates the operations which are
// required to fulfill the instruction, executing these against the DB. It
// also returns the result of the instruction.
func (e *executor) execute(instr instruction) (result, error) {
	switch instr.command {
	case commandInsert:
		return result{}, fmt.Errorf("unimplemented")
	case commandSelect:
		return e.executeSelect(instr)
	case commandDelete:
		return result{}, fmt.Errorf("unimplemented")
	case commandCreateTable:
		return e.executeCreateTable(instr)

	default:
		return result{}, fmt.Errorf("invalid executor command")
	}
}

// Executes the select query instruction, returning the structure of the table
// (columns) and the rows specified in the query.
func (e *executor) executeSelect(instr instruction) (result, error) {
	_, exists := e.db.tables[instr.table]
	if !exists {
		return result{}, fmt.Errorf("table %s does not exist", instr.table)
	}

	// TODO check if columns all exist in table
	// TODO btree get all method

	return result{}, fmt.Errorf("unimplemented")
}

// Executes the create table instruction, parses the columns given as arguments
// and adds a new table record to the storage map.
func (e *executor) executeCreateTable(instr instruction) (result, error) {
	cols, err := parseInsertColumns(instr.params)
	if err != nil {
		return result{}, fmt.Errorf("failed to parse column params: %v", err)
	}

	e.db.tables[instr.table] = table{
		name:    instr.table,
		store:   btree.NewBtreeOrder(e.cfg.order),
		columns: cols,
	}

	return result{created: 1}, nil
}

func parseInsertColumns(params []string) ([]column, error) {
	// If there are no tables to be created, return early
	if len(params) == 0 {
		return []column{}, nil
	}

	// There should be a column name, type and nullable field for each column
	// which is being declared
	if len(params)%3 != 0 {
		return []column{}, fmt.Errorf("invalid column pairs, every name must have a type")
	}

	cols := make([]column, 0, len(params)/3)

	for i := 0; i < len(params); i += 3 {
		p1, p2, p3 := params[i], params[i+1], params[i+2]
		if err := validateTableName(p1); err != nil {
			return cols, fmt.Errorf("invalid table name: %s: %v", p1, err)
		}

		colType := parseColumnType(p2)
		if colType == columnTypeInvalid {
			return cols, fmt.Errorf("found invalid column type: %s", p2)
		}

		b, err := parseBool(p3)
		if err != nil {
			return cols, fmt.Errorf("invalid value for field is_nullable: %s", err)
		}

		cols = append(cols, column{
			name:       p1,
			dataType:   colType,
			isNullable: b,
		})
	}

	return cols, nil
}

// The maximum number of characters allowed in a table name
const tableNameMaxLen = 32

// The validation pattern used to determine whether a table name is valid
var tableNamePattern = regexp.MustCompile(`^[a-zA-Z]+$`)

// Validates whether the string can be used as a valid table name identifier.
// Returns an error if it is invalid, nil if valid.
func validateTableName(name string) error {
	if len(name) > tableNameMaxLen {
		return fmt.Errorf("table name exceeds the character limit of %d", tableNameMaxLen)
	}

	if !tableNamePattern.MatchString(name) {
		return fmt.Errorf("table name includes invalid characters")
	}

	return nil
}

func parseBool(str string) (bool, error) {
	switch str {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean field")
	}
}
