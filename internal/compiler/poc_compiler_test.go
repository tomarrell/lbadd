package compiler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/parser"
)

func Test_pocCompiler_Compile(t *testing.T) {
	tests := []struct {
		name    string
		stmt    string
		want    command.Command
		wantErr bool
	}{
		{
			"select simple",
			"SELECT col FROM users WHERE otherColumn IS someValue",
			command.Select{
				Tables: []string{"users"},
				Cols:   []string{"col"},
				Where:  nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			c := &pocCompiler{}
			p := parser.New(tt.stmt)
			stmt, errs, ok := p.Next()
			assert.Nil(errs)
			for _, err := range errs {
				assert.NoError(err)
			}
			assert.True(ok)

			got, err := c.Compile(stmt)
			assert.Equal(tt.want, got)
			if tt.wantErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
		})
	}
}
