package engine

import (
	"sync"

	"github.com/tomarrell/lbadd/internal/id"
)

// ExecutionContext is a context that is passed down throughout a complete
// evaluation. It may be populated further.
type ExecutionContext struct {
	*executionContext
}

type executionContext struct {
	id                id.ID
	scannedTablesLock sync.Mutex
	scannedTables     map[string]Table
}

func newEmptyExecutionContext() ExecutionContext {
	return ExecutionContext{
		executionContext: &executionContext{
			id:            id.Create(),
			scannedTables: make(map[string]Table),
		},
	}
}

func (c ExecutionContext) putScannedTable(name string, table Table) {
	c.scannedTablesLock.Lock()
	defer c.scannedTablesLock.Unlock()

	c.scannedTables[name] = table
}

func (c ExecutionContext) getScannedTable(name string) (Table, bool) {
	c.scannedTablesLock.Lock()
	defer c.scannedTablesLock.Unlock()

	tbl, ok := c.scannedTables[name]
	return tbl, ok
}

func (c ExecutionContext) String() string {
	return c.id.String()
}
