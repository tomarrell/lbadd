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
			"simple values",
			"VALUES (1,2,3),(4,5,6),(7,8,9)",
			command.Values{
				Values: [][]command.Expr{
					{
						command.LiteralExpr{Value: "1"},
						command.LiteralExpr{Value: "2"},
						command.LiteralExpr{Value: "3"},
					},
					{
						command.LiteralExpr{Value: "4"},
						command.LiteralExpr{Value: "5"},
						command.LiteralExpr{Value: "6"},
					},
					{
						command.LiteralExpr{Value: "7"},
						command.LiteralExpr{Value: "8"},
						command.LiteralExpr{Value: "9"},
					},
				},
			},
			false,
		},
		{
			"simple select",
			"SELECT * FROM myTable",
			command.Project{
				Cols: []command.Column{
					{
						Column: command.LiteralExpr{Value: "*"},
					},
				},
				Input: command.Scan{
					Table: command.SimpleTable{
						Table: "myTable",
					},
				},
			},
			false,
		},
		{
			"simple select where",
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
			"simple select limit",
			"SELECT * FROM myTable LIMIT 5",
			command.Limit{
				Limit: command.LiteralExpr{Value: "5"},
				Input: command.Project{
					Cols: []command.Column{
						{
							Column: command.LiteralExpr{Value: "*"},
						},
					},
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
			"simple select limit offset",
			"SELECT * FROM myTable LIMIT 5, 10",
			command.Limit{
				Limit: command.LiteralExpr{Value: "5"},
				Input: command.Offset{
					Offset: command.LiteralExpr{Value: "10"},
					Input: command.Project{
						Cols: []command.Column{
							{
								Column: command.LiteralExpr{Value: "*"},
							},
						},
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
			"simple select limit offset",
			"SELECT * FROM myTable LIMIT 5 OFFSET 10",
			command.Limit{
				Limit: command.LiteralExpr{Value: "5"},
				Input: command.Offset{
					Offset: command.LiteralExpr{Value: "10"},
					Input: command.Project{
						Cols: []command.Column{
							{
								Column: command.LiteralExpr{Value: "*"},
							},
						},
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
		{
			"select expression",
			"SELECT name, amount * price AS total_price FROM items JOIN prices",
			command.Project{
				Cols: []command.Column{
					{
						Column: command.LiteralExpr{Value: "name"},
					},
					{
						Column: command.BinaryExpr{
							Operator: "*",
							Left:     command.LiteralExpr{Value: "amount"},
							Right:    command.LiteralExpr{Value: "price"},
						},
						Alias: "total_price",
					},
				},
				Input: command.Join{
					Left: command.Scan{
						Table: command.SimpleTable{Table: "items"},
					},
					Right: command.Scan{
						Table: command.SimpleTable{Table: "prices"},
					},
				},
			},
			false,
		},
		{
			"select function",
			"SELECT AVG(price) AS avg_price FROM items LEFT JOIN prices",
			command.Project{
				Cols: []command.Column{
					{
						Column: command.FunctionExpr{
							Name:     "AVG",
							Distinct: false,
							Args: []command.Expr{
								command.LiteralExpr{Value: "price"},
							},
						},
						Alias: "avg_price",
					},
				},
				Input: command.Join{
					Type: command.JoinLeft,
					Left: command.Scan{
						Table: command.SimpleTable{Table: "items"},
					},
					Right: command.Scan{
						Table: command.SimpleTable{Table: "prices"},
					},
				},
			},
			false,
		},
		{
			"select function distinct",
			"SELECT AVG(DISTINCT price) AS avg_price FROM items LEFT JOIN prices",
			command.Project{
				Cols: []command.Column{
					{
						Column: command.FunctionExpr{
							Name:     "AVG",
							Distinct: true,
							Args: []command.Expr{
								command.LiteralExpr{Value: "price"},
							},
						},
						Alias: "avg_price",
					},
				},
				Input: command.Join{
					Type: command.JoinLeft,
					Left: command.Scan{
						Table: command.SimpleTable{Table: "items"},
					},
					Right: command.Scan{
						Table: command.SimpleTable{Table: "prices"},
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
			assert.Len(errs, 0)
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
