package driver

import (
	"context"
	"database/sql/driver"
	"fmt"
	"strconv"
)

var _ driver.Stmt = (*Stmt)(nil)
var _ driver.StmtExecContext = (*Stmt)(nil)
var _ driver.StmtQueryContext = (*Stmt)(nil)

type Stmt struct {
}

func parse(query string) (*Stmt, error) {
	return nil, fmt.Errorf("unimplemented") // TODO(TimSatke): implement
}

func (s *Stmt) Close() error {
	return nil // TODO(TimSatke): implement
}

func (s *Stmt) NumInput() int {
	return 0 // TODO(TimSatke): implement
}

func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	namedValues := make([]driver.NamedValue, 0, len(args))
	for i, arg := range args {
		namedValues = append(namedValues, driver.NamedValue{
			Name:    "param" + strconv.Itoa(i),
			Ordinal: i,
			Value:   arg,
		})
	}
	return s.ExecContext(context.Background(), namedValues)
}

func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	namedValues := make([]driver.NamedValue, 0, len(args))
	for i, arg := range args {
		namedValues = append(namedValues, driver.NamedValue{
			Name:    "param" + strconv.Itoa(i),
			Ordinal: i,
			Value:   arg,
		})
	}
	return s.QueryContext(context.Background(), namedValues)
}

func (s *Stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return nil, fmt.Errorf("unimplemented") // TODO(TimSatke): implement
}

func (s *Stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return nil, fmt.Errorf("unimplemented") // TODO(TimSatke): implement
}
