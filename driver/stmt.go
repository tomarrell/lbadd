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

// Stmt is a prepared statement that can be executed. It does not remember
// values that were passed in.
type Stmt struct {
}

// parse attempts to parse the given query string. If the query string is valid
// and supported sql, a statement and error=nil will be returned.
func parse(query string) (*Stmt, error) {
	return nil, fmt.Errorf("unimplemented") // TODO(TimSatke): implement
}

// Close closes this statement, making it impossible to execute it again.
func (s *Stmt) Close() error {
	return nil // TODO(TimSatke): implement
}

// NumInput returns the amount of argument placeholders that the statement has.
func (s *Stmt) NumInput() int {
	return 0 // TODO(TimSatke): implement
}

// Exec is discouraged. Don't use this, use ExecContext instead.
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

// Query is discouraged. Don't use this, use QueryContext instead.
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

// ExecContext executes this statement with the given arguments as arguments,
// with respect to the given context. This should be used for update statements only (alter, update, drop, delete etc.).
func (s *Stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return nil, fmt.Errorf("unimplemented") // TODO(TimSatke): implement
}

// QueryContext executes this statement with the given arguments as arguments,
// with respect to the given context. This should be used for query statements
// only (select etc.).
func (s *Stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return nil, fmt.Errorf("unimplemented") // TODO(TimSatke): implement
}
