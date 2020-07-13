package test

import "testing"

func TestIssue188(t *testing.T) {
	RunAndCompare(t, Test{
		Name:      "issue188",
		Statement: `VALUES ("abc", "a\"bc", "a\u0062c")`,
	})
}
