package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCommand(t *testing.T) {
	cases := []struct {
		name     string
		command  string
		expected instruction
	}{
		{
			name:     "insert command",
			command:  "insert users a b",
			expected: instruction{commandInsert, "users", []string{"a", "b"}},
		},
		{
			name:     "select command",
			command:  "select table a b c",
			expected: instruction{commandSelect, "table", []string{"a", "b", "c"}},
		},
		{
			name:     "select with filter",
			command:  "select table a b c<1",
			expected: instruction{commandSelect, "table", []string{"a", "b", "c<1"}},
		},
		{
			name:     "delete command",
			command:  "delete table a>6 b=1",
			expected: instruction{commandDelete, "table", []string{"a>6", "b=1"}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			repl := newRepl()
			instr, err := repl.readCommand(tc.command)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, instr)
		})
	}
}
