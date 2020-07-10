package engine

import (
	"sync"

	"github.com/tomarrell/lbadd/internal/id"
)

// ExecutionContext is a context that is passed down throughout a complete
// evaluation. It may be populated further.
type ExecutionContext struct {
	ctx *executionContext
}

type executionContext struct {
	id                id.ID
	scannedTablesLock sync.Mutex
	scannedTables     map[string]Table
}

func newEmptyExecutionContext() ExecutionContext {
	return ExecutionContext{
		ctx: &executionContext{
			id:            id.Create(),
			scannedTables: make(map[string]Table),
		},
	}
}

func (c ExecutionContext) putScannedTable(name string, table Table) {
	c.ctx.scannedTablesLock.Lock()
	defer c.ctx.scannedTablesLock.Unlock()

	c.ctx.scannedTables[name] = table
}

func (c ExecutionContext) getScannedTable(name string) (Table, bool) {
	c.ctx.scannedTablesLock.Lock()
	defer c.ctx.scannedTablesLock.Unlock()

	tbl, ok := c.ctx.scannedTables[name]
	return tbl, ok
}

func (c ExecutionContext) String() string {
	return c.ctx.id.String()
}
