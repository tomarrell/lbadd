package executor

import "fmt"

// Result describes the result of a command execution. The result is always a
// table that has a header row. The smallest possible result table is a table
// with one column and two rows, and is generated as a result of a single-value
// computation, e.g. a sum().
type Result interface {
	fmt.Stringer
}
