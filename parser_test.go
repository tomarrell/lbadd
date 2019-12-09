package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// queryType  queryType
// tableName  string
// conditions []condition
// updates    map[string]string
// inserts    [][]string
// fields     []string

func TestParser(t *testing.T) {
	cases := []struct {
		name     string
		sql      string
		expected query
		err      error
	}{
		{
			name:     "select single field from table",
			sql:      "SELECT a FROM z",
			expected: query{queryType: selectQuery, fields: []string{"a"}, tableName: "z"},
		},
		{
			name:     "select multiple fields from table",
			sql:      "SELECT a, b, c FROM z",
			expected: query{queryType: selectQuery, fields: []string{"a", "b", "c"}, tableName: "z"},
		},
		{
			name:     "select with field and trailing comma error",
			sql:      "SELECT a, b, c FROM z",
			expected: query{queryType: selectQuery, fields: []string{"a", "b", "c"}, tableName: "z"},
			err:      fmt.Errorf("invalid trailing comma"),
		},
		{
			name:     "select all (*) fields from table",
			sql:      "SELECT * FROM z",
			expected: query{queryType: selectQuery, fields: []string{"*"}, tableName: "z"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := parse(tc.sql)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
