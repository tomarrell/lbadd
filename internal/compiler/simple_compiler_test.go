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
						Column: command.LiteralExpr{Value: "*"},
					},
				},
				Input: command.Select{
					Filter: command.LiteralExpr{Value: "true"},
					Input: command.Scan{
						Table: command.SimpleTable{
							Table: "myTable",
						},
					},
				},
			},
			false,
		},
		{
			"select distinct",
			"SELECT DISTINCT * FROM myTable WHERE true",
			command.Distinct{
				Input: command.Project{
					Cols: []command.Column{
						{
							Column: command.LiteralExpr{Value: "*"},
						},
					},
					Input: command.Select{
						Filter: command.LiteralExpr{Value: "true"},
						Input: command.Scan{
							Table: command.SimpleTable{
								Table: "myTable",
							},
						},
					},
				},
			},
			false,
		},
		{
			"select with implicit join",
			"SELECT * FROM a, b WHERE true",
			command.Project{
				Cols: []command.Column{
					{
						Column: command.LiteralExpr{Value: "*"},
					},
				},
				Input: command.Select{
					Filter: command.LiteralExpr{Value: "true"},
					Input: command.Join{
						Left: command.Scan{
							Table: command.SimpleTable{
								Table: "a",
							},
						},
						Right: command.Scan{
							Table: command.SimpleTable{
								Table: "b",
							},
						},
					},
				},
			},
			false,
		},
		{
			"select with explicit join",
			"SELECT * FROM a JOIN b WHERE true",
			command.Project{
				Cols: []command.Column{
					{
						Column: command.LiteralExpr{Value: "*"},
					},
				},
				Input: command.Select{
					Filter: command.LiteralExpr{Value: "true"},
					Input: command.Join{
						Left: command.Scan{
							Table: command.SimpleTable{
								Table: "a",
							},
						},
						Right: command.Scan{
							Table: command.SimpleTable{
								Table: "b",
							},
						},
					},
				},
			},
			false,
		},
		{
			"select with implicit and explicit join",
			"SELECT * FROM a, b JOIN c WHERE true",
			command.Project{
				Cols: []command.Column{
					{
						Column: command.LiteralExpr{Value: "*"},
					},
				},
				Input: command.Select{
					Filter: command.LiteralExpr{Value: "true"},
					Input: command.Join{
						Left: command.Join{
							Left: command.Scan{
								Table: command.SimpleTable{
									Table: "a",
								},
							},
							Right: command.Scan{
								Table: command.SimpleTable{
									Table: "b",
								},
							},
						},
						Right: command.Scan{
							Table: command.SimpleTable{
								Table: "c",
							},
						},
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

			if tt.wantErr {
				assert.Error(gotErr)
			} else {
				assert.NoError(gotErr)
			}
			assert.Equal(tt.want, got)
		})
	}
}
