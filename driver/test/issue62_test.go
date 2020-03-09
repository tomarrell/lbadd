package test

import (
	"database/sql"
	"testing"
)

func TestIssue62(t *testing.T) {
	/*
		This test does not do anything. This is an example to show how
		issue-based integration tests are written. This package contains a
		test-main, which sets up and tears down a database server. In the tests,
		obtain a connection pool and execute statements; in general, try to
		reproduce the error that is described in the issue.

		This generally works best if the issue was a reproducible sql statement,
		that produced an error inside the parser/executor/db or wherever.

		Assume that there is a database up and running on the address that is
		stored inside TestAddress.
	*/
	pool, _ := sql.Open("lbadd", TestAddress)
	if pool != nil {
		_ = pool.Close()
	}
}
