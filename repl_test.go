package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCommand(t *testing.T) {
	cases := []struct {
		name     string
		command  string
		expected instruction
	}{
		{"insert command", "insert table a b", instruction{commandInsert}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			repl := newRepl()
			instr, err := repl.readCommand(strings.NewReader(tc.command))
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, instr)
		})
	}
}
