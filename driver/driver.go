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

// Driver is the database driver that can communicate with an lbadd database. It
// will be registered with the name "lbadd".
type Driver struct {
}

// Open creates a new connector and uses that connector to open a new
// connection. The context that is used to open the new connection is
// context.Background().
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

// OpenConnector creates a connector that can be used to open a connection to a
// data source. The data source is specified by the given name. A connector can
// only open connections to his data source.
func (d *Driver) OpenConnector(name string) (driver.Connector, error) {
	return &Connector{
		driver: d,
	}, nil
}
