package driver

import (
	"context"
	"database/sql/driver"
)

var _ driver.Connector = (*Connector)(nil)

type Connector struct {
	driver *Driver
}

func (c *Connector) Connect(ctx context.Context) (driver.Conn, error) {
	cancelableCtx, cancel := context.WithCancel(ctx)
	return &Conn{
		ctx:    cancelableCtx,
		cancel: cancel,
	}, nil
}

func (c *Connector) Driver() driver.Driver {
	return c.driver
}
