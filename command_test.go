package lbadd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newCommand(t *testing.T) {
	type args struct {
		cmd string
	}

	tests := []struct {
		name string
		args args
		want command
	}{
		{
			name: "unknown command",
			args: args{cmd: "uh oh"},
			want: 0,
		},
		{
			name: "insert",
			args: args{cmd: "insert"},
			want: 1,
		},
		{
			name: "select",
			args: args{cmd: "select"},
			want: 2,
		},
		{
			name: "delete",
			args: args{cmd: "delete"},
			want: 3,
		},
		{
			name: "mixed casing insert",
			args: args{cmd: "iNsErT"},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newCommand(tt.args.cmd)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_command_String(t *testing.T) {
	tests := []struct {
		name string
		c    command
		want string
	}{
		{
			name: "unknown",
			c:    0,
			want: "UNKNOWN",
		},
		{
			name: "insert",
			c:    1,
			want: "INSERT",
		},
		{
			name: "select",
			c:    2,
			want: "SELECT",
		},
		{
			name: "delete",
			c:    3,
			want: "DELETE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
