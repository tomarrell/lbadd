package driver

import (
	"context"
	"database/sql/driver"
)

var _ driver.Conn = (*Conn)(nil)
var _ driver.ConnBeginTx = (*Conn)(nil)

type Conn struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	if c.ctx.Err() != nil {
		return nil, ErrConnectionClosed
	}

	return nil, nil // TODO(TimSatke) implement
}

func (c *Conn) Close() error {
	c.cancel()
	return nil
}

func (c *Conn) Begin() (driver.Tx, error) {
	return c.BeginTx(c.ctx, driver.TxOptions{})
}

func (c *Conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return nil, nil // TODO(TimSatke) implement
}
