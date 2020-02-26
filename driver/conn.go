package driver

import (
	"context"
	"database/sql/driver"
)

var _ driver.Conn = (*Conn)(nil)
var _ driver.ConnBeginTx = (*Conn)(nil)
var _ driver.ExecerContext = (*Conn)(nil)
var _ driver.Pinger = (*Conn)(nil)
var _ driver.QueryerContext = (*Conn)(nil)

// Conn represents a connection to the database. It can be used to prepare and
// execute statements.
type Conn struct {
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return nil, nil // TODO(TimSatke): implement
}

func (c *Conn) Close() error {
	return nil // TODO(TimSatke): implement
}

func (c *Conn) Begin() (driver.Tx, error) {
	return nil, nil // TODO(TimSatke): implement
}

func (c *Conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return nil, nil // TODO(TimSatke): implement
}

func (c *Conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	return nil, nil // TODO(TimSatke): implement
}

func (c *Conn) Ping(ctx context.Context) error {
	return nil // TODO(TimSatke): implement
}

func (c *Conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	return nil, nil // TODO(TimSatke): implement
}
