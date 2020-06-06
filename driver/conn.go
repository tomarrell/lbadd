package driver

import (
	"context"
	"database/sql/driver"
	"fmt"
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

// Prepare prepares a statement. The returned Stmt is an SQL prepared statement,
// that has placeholders for parameters that need to be set. Do NOT set
// parameters directly when creating the statement with this method.
//
//  stmt, err := conn.Prepare(`INSERT INTO users VALUES ("jdoe")`) // WRONG
//
//  stmt, err := conn.Prepare(`INSERT INTO users VALUES (?)`) // CORRECT
//  result, err := stmt.Exec("jdoe")
func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	stmt, err := parse(query)
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}
	return stmt, nil
}

// Close closes this connection. If a connection is closed, it cannot be used as
// idle connection in the connection pool and a new connection needs to be
// established.
func (c *Conn) Close() error {
	return nil // TODO(TimSatke): implement
}

// Begin is deprecated. Use BeginTx instead.
func (c *Conn) Begin() (driver.Tx, error) {
	return c.BeginTx(context.Background(), driver.TxOptions{})
}

// BeginTx creates a transaction that can be either committed or rolled back.
func (c *Conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return nil, fmt.Errorf("unimplemented") // TODO(TimSatke): implement
}

// ExecContext executes the given query with the given arguments under the given
// context and returns an exec result. The statement must contain placeholders,
// one for each element of the given arguments.
func (c *Conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	rawStmt, err := c.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare: %w", err)
	}

	if stmt, ok := rawStmt.(*Stmt); ok {
		return stmt.ExecContext(ctx, args)
	}

	return nil, fmt.Errorf("cannot execute Stmt of type %T, expected %T", rawStmt, &Stmt{})
}

// Ping pings the database, failing if the connection is closed or the database
// failed.
func (c *Conn) Ping(ctx context.Context) error {
	return nil // TODO(TimSatke): implement
}

// QueryContext executes the given query with the given arguments under the
// given context and returns a query result. The query must contain
// placeholders, one for each element of the given arguments.
func (c *Conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	rawStmt, err := c.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare: %w", err)
	}

	if stmt, ok := rawStmt.(*Stmt); ok {
		return stmt.QueryContext(ctx, args)
	}

	return nil, fmt.Errorf("cannot execute query Stmt of type %T, expected %T", rawStmt, &Stmt{})
}
