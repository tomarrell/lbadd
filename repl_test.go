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
		{"insert command", "insert table a b", instruction{commandInsert, []string{"table", "a", "b"}}},
		{"select command", "select table a b c", instruction{commandSelect, []string{"table", "a", "b", "c"}}},
		{"delete command", "delete table a>6 b=1", instruction{commandDelete, []string{"table", "a>6", "b=1"}}},
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
