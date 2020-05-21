package compiler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/parser"
)

func Test_simpleCompiler_Compile_NoOptimizations(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    command.Command
		wantErr bool
	}{
		{
			"simple select",
			"SELECT * FROM myTable WHERE true",
			command.Project{
				Cols: []command.Column{
					{
						Table:  "",
						Column: command.LiteralExpr{Value: "*"},
						Alias:  "",
					},
				},
				Input: command.Select{
					Filter: command.LiteralExpr{Value: "true"},
					Input: command.Join{
						Filter: nil,
						Left: command.Scan{
							Table: command.SimpleTable{
								Schema:  "",
								Table:   "myTable",
								Alias:   "",
								Indexed: true,
								Index:   "",
							},
						},
						Right: nil,
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			c := &simpleCompiler{}
			p := parser.New(tt.input)
			stmt, errs, ok := p.Next()
			assert.Nil(errs)
			assert.True(ok)

			got, gotErr := c.Compile(stmt)

			assert.Equal(tt.want, got)
			if tt.wantErr {
				assert.Error(gotErr)
			} else {
				assert.NoError(gotErr)
			}
		})
	}
}
