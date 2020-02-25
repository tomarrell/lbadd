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
	return nil, nil // TODO(TimSatke): implement
}

func (c *Connector) Driver() driver.Driver {
	return c.driver
}
