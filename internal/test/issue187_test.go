package test

import "testing"

func TestIssue187(t *testing.T) {
	RunAndCompare(t, Test{
		Name:      "issue187",
		Statement: `VALUES (1,"2",3), (4,"5",6)`,
	})
}
