package driver

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
)

func init() {
	sql.Register("lbadd", &Driver{})
}

var _ driver.Driver = (*Driver)(nil)
var _ driver.DriverContext = (*Driver)(nil)

type Driver struct {
}

func (d *Driver) Open(name string) (driver.Conn, error) {
	connector, err := d.OpenConnector(name)
	if err != nil {
		return nil, fmt.Errorf("open connector: %w", err)
	}

	conn, err := connector.Connect(context.Background())
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}
	return conn, nil
}

func (d *Driver) OpenConnector(name string) (driver.Connector, error) {
	return &Connector{
		driver: d,
	}, nil
}
