package driver

import (
	"context"
	"database/sql/driver"
)

var _ driver.Stmt = (*Stmt)(nil)
var _ driver.StmtExecContext = (*Stmt)(nil)
var _ driver.StmtQueryContext = (*Stmt)(nil)

type Stmt struct {
}

func (s *Stmt) Close() error {
	return nil // TODO(TimSatke): implement
}

func (s *Stmt) NumInput() int {
	return 0 // TODO(TimSatke): implement
}

func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, nil // TODO(TimSatke): implement
}

func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	return nil, nil // TODO(TimSatke): implement
}

func (s *Stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return nil, nil // TODO(TimSatke): implement
}

func (s *Stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return nil, nil // TODO(TimSatke): implement
}
