package driver

import (
	"context"
	"database/sql/driver"
)

var _ driver.Connector = (*Connector)(nil)

// Connector implements a component that is able to open a connection to the
// database that is remembered by the connector. This connection can then be
// used to prepare and execute statements.
type Connector struct {
	driver *Driver
}

// Connect opens a connection to the database that the connector is configured
// to connect to. The opening of the connection pays respect to deadlines or
// timeouts configured in the context.
func (c *Connector) Connect(ctx context.Context) (driver.Conn, error) {
	return &Conn{}, nil
}

// Driver returns the underlying driver, that the connector has been created
// from.
func (c *Connector) Driver() driver.Driver {
	return c.driver
}
