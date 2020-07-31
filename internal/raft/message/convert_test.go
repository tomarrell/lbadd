package message

import (
	"reflect"
	"testing"

	"github.com/tomarrell/lbadd/internal/compiler/command"
)

var commandToMessageTests = []struct {
	in  command.Command
	out Message
}{
	{
		// SCAN
		&command.Scan{
			Table: &command.SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
		},
		&Command_Scan{
			Table: &SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
		},
	},
	{
		// SELECT
		&command.Select{
			Filter: &command.LiteralExpr{
				Value: "literal",
			},
			Input: &command.Scan{
				Table: &command.SimpleTable{
					Schema:  "mySchema",
					Table:   "myTable",
					Alias:   "myAlias",
					Indexed: true,
					Index:   "myIndex",
				},
			},
		},
		&Command_Select{
			Filter: &Expr{
				Expr: &Expr_Literal{
					&LiteralExpr{
						Value: "literal",
					},
				},
			},
			Input: &List{
				List: &List_Scan{
					Scan: &Command_Scan{
						Table: &SimpleTable{
							Schema:  "mySchema",
							Table:   "myTable",
							Alias:   "myAlias",
							Indexed: true,
							Index:   "myIndex",
						},
					},
				},
			},
		},
	},
	{
		// PROJECT
		&command.Project{
			Cols: []command.Column{
				{
					Table: "myTable1",
					Column: &command.LiteralExpr{
						Value: "literal",
					},
					Alias: "myAlias1",
				},
				{
					Table: "myTable2",
					Column: &command.LiteralExpr{
						Value: "literal",
					},
					Alias: "myAlias2",
				},
			},
			Input: &command.Scan{
				Table: &command.SimpleTable{
					Schema:  "mySchema",
					Table:   "myTable",
					Alias:   "myAlias",
					Indexed: true,
					Index:   "myIndex",
				},
			},
		},
		&Command_Project{
			Cols: []*Column{
				{
					Table: "myTable1",
					Column: &Expr{
						Expr: &Expr_Literal{
							&LiteralExpr{
								Value: "literal",
							},
						},
					},
					Alias: "myAlias1",
				},
				{
					Table: "myTable2",
					Column: &Expr{
						Expr: &Expr_Literal{
							&LiteralExpr{
								Value: "literal",
							},
						},
					},
					Alias: "myAlias2",
				},
			},
			Input: &List{
				List: &List_Scan{
					Scan: &Command_Scan{
						Table: &SimpleTable{
							Schema:  "mySchema",
							Table:   "myTable",
							Alias:   "myAlias",
							Indexed: true,
							Index:   "myIndex",
						},
					},
				},
			},
		},
	},
	{
		// DELETE
		&command.Delete{
			Table: &command.SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
			Filter: &command.BinaryExpr{
				Operator: "operator",
				Left: &command.LiteralExpr{
					Value: "leftLiteral",
				},
				Right: &command.LiteralExpr{
					Value: "rightLiteral",
				},
			},
		},
		&Command_Delete{
			Table: &SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
			Filter: &Expr{
				Expr: &Expr_Binary{
					Binary: &BinaryExpr{
						Operator: "operator",
						Left: &Expr{
							Expr: &Expr_Literal{
								Literal: &LiteralExpr{
									Value: "leftLiteral",
								},
							},
						},
						Right: &Expr{
							Expr: &Expr_Literal{
								Literal: &LiteralExpr{
									Value: "rightLiteral",
								},
							},
						},
					},
				},
			},
		},
	},
	{
		// DROP TABLE
		&command.DropTable{
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
		&CommandDrop{
			Target:   0,
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
	},
	{
		// DROP VIEW
		&command.DropView{
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
		&CommandDrop{
			Target:   1,
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
	},
	{
		// DROP INDEX
		&command.DropIndex{
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
		&CommandDrop{
			Target:   2,
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
	},
	{
		// DROP TRIGGER
		&command.DropTrigger{
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
		&CommandDrop{
			Target:   3,
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
	},
	{
		// UPDATE
		&command.Update{
			UpdateOr: 0,
			Table: &command.SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
			Updates: []command.UpdateSetter{
				{
					Cols: []string{
						"col1",
						"col2",
					},
					Value: command.ConstantBooleanExpr{
						Value: true,
					},
				},
			},
			Filter: &command.EqualityExpr{
				Left: &command.LiteralExpr{
					Value: "leftLiteral",
				},
				Right: &command.LiteralExpr{
					Value: "rightLiteral",
				},
			},
		},
		&Command_Update{
			UpdateOr: 0,
			Table: &SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
			Updates: []*UpdateSetter{
				{
					Cols: []string{
						"col1",
						"col2",
					},
					Value: &UpdateSetter_Constant{
						Constant: &ConstantBooleanExpr{
							Value: true,
						},
					},
				},
			},
			Filter: &Expr{
				Expr: &Expr_Equality{
					Equality: &EqualityExpr{
						Left: &Expr{
							Expr: &Expr_Literal{
								Literal: &LiteralExpr{
									Value: "leftLiteral",
								},
							},
						},
						Right: &Expr{
							Expr: &Expr_Literal{
								Literal: &LiteralExpr{
									Value: "rightLiteral",
								},
							},
						},
					},
				},
			},
		},
	},
	{
		// JOIN
		&command.Join{
			Natural: true,
			Type:    0,
			Filter: &command.FunctionExpr{
				Name:     "function",
				Distinct: true,
				Args: []command.Expr{
					&command.RangeExpr{
						Needle: &command.LiteralExpr{
							Value: "literal",
						},
						Lo: &command.LiteralExpr{
							Value: "literal",
						},
						Hi: &command.LiteralExpr{
							Value: "literal",
						},
						Invert: false,
					},
					&command.UnaryExpr{
						Operator: "operator",
						Value: &command.LiteralExpr{
							Value: "literal",
						},
					},
				},
			},
			Left: &command.Scan{
				Table: &command.SimpleTable{
					Schema:  "mySchema",
					Table:   "myTable",
					Alias:   "myAlias",
					Indexed: true,
					Index:   "myIndex",
				},
			},
			Right: &command.Scan{
				Table: &command.SimpleTable{
					Schema:  "mySchema",
					Table:   "myTable",
					Alias:   "myAlias",
					Indexed: true,
					Index:   "myIndex",
				},
			},
		},
		&Command_Join{
			Natural: true,
			Type:    0,
			Filter: &Expr{
				Expr: &Expr_Func{
					Func: &FunctionExpr{
						Name:     "function",
						Distinct: true,
						Args: []*Expr{
							{
								Expr: &Expr_Range{
									&RangeExpr{
										Needle: &Expr{
											Expr: &Expr_Literal{
												Literal: &LiteralExpr{
													Value: "literal",
												},
											},
										},
										Lo: &Expr{
											Expr: &Expr_Literal{
												Literal: &LiteralExpr{
													Value: "literal",
												},
											},
										},
										Hi: &Expr{
											Expr: &Expr_Literal{
												Literal: &LiteralExpr{
													Value: "literal",
												},
											},
										},
										Invert: false,
									},
								},
							},
							{
								Expr: &Expr_Unary{
									Unary: &UnaryExpr{
										Operator: "operator",
										Value: &Expr{
											Expr: &Expr_Literal{
												Literal: &LiteralExpr{
													Value: "literal",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			Left: &List{
				List: &List_Scan{
					Scan: &Command_Scan{
						Table: &SimpleTable{
							Schema:  "mySchema",
							Table:   "myTable",
							Alias:   "myAlias",
							Indexed: true,
							Index:   "myIndex",
						},
					},
				},
			},
			Right: &List{
				List: &List_Scan{
					Scan: &Command_Scan{
						Table: &SimpleTable{
							Schema:  "mySchema",
							Table:   "myTable",
							Alias:   "myAlias",
							Indexed: true,
							Index:   "myIndex",
						},
					},
				},
			},
		},
	},
	{
		// LIMIT
		&command.Limit{
			Limit: &command.LiteralExpr{
				Value: "literal",
			},
			Input: &command.Scan{
				Table: &command.SimpleTable{
					Schema:  "mySchema",
					Table:   "myTable",
					Alias:   "myAlias",
					Indexed: true,
					Index:   "myIndex",
				},
			},
		},
		&Command_Limit{
			Limit: &Expr{
				Expr: &Expr_Literal{
					Literal: &LiteralExpr{
						Value: "literal",
					},
				},
			},
			Input: &List{
				List: &List_Scan{
					Scan: &Command_Scan{
						Table: &SimpleTable{
							Schema:  "mySchema",
							Table:   "myTable",
							Alias:   "myAlias",
							Indexed: true,
							Index:   "myIndex",
						},
					},
				},
			},
		},
	},
	{
		// INSERT
		&command.Insert{
			InsertOr: 0,
			Table: &command.SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
			Cols: []command.Column{
				{
					Table: "myTable1",
					Column: &command.LiteralExpr{
						Value: "literal",
					},
					Alias: "myAlias1",
				},
				{
					Table: "myTable2",
					Column: &command.LiteralExpr{
						Value: "literal",
					},
					Alias: "myAlias2",
				},
			},
			DefaultValues: false,
			Input: &command.Scan{
				Table: &command.SimpleTable{
					Schema:  "mySchema",
					Table:   "myTable",
					Alias:   "myAlias",
					Indexed: true,
					Index:   "myIndex",
				},
			},
		},
		&Command_Insert{
			InsertOr: 0,
			Table: &SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
			Cols: []*Column{
				{
					Table: "myTable1",
					Column: &Expr{
						Expr: &Expr_Literal{
							&LiteralExpr{
								Value: "literal",
							},
						},
					},
					Alias: "myAlias1",
				},
				{
					Table: "myTable2",
					Column: &Expr{
						Expr: &Expr_Literal{
							&LiteralExpr{
								Value: "literal",
							},
						},
					},
					Alias: "myAlias2",
				},
			},
			DefaultValues: false,
			Input: &List{
				List: &List_Scan{
					Scan: &Command_Scan{
						Table: &SimpleTable{
							Schema:  "mySchema",
							Table:   "myTable",
							Alias:   "myAlias",
							Indexed: true,
							Index:   "myIndex",
						},
					},
				},
			},
		},
	},
}

func TestConvertCommandToMessage(t *testing.T) {
	for _, tt := range commandToMessageTests {
		t.Run(tt.in.String(), func(t *testing.T) {
			msg, _ := ConvertCommandToMessage(tt.in)
			if !reflect.DeepEqual(msg, tt.out) {
				t.Errorf("got %q, want %q", msg, tt.out)
			}
		})
	}
}

var messageToCommandTests = []struct {
	in  Message
	out command.Command
}{
	{
		// SCAN
		&Command_Scan{
			Table: &SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
		},
		&command.Scan{
			Table: &command.SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
		},
	},
	{
		// SELECT
		&Command_Select{
			Filter: &Expr{
				Expr: &Expr_Literal{
					&LiteralExpr{
						Value: "literal",
					},
				},
			},
			Input: &List{
				List: &List_Scan{
					Scan: &Command_Scan{
						Table: &SimpleTable{
							Schema:  "mySchema",
							Table:   "myTable",
							Alias:   "myAlias",
							Indexed: true,
							Index:   "myIndex",
						},
					},
				},
			},
		},
		&command.Select{
			Filter: &command.LiteralExpr{
				Value: "literal",
			},
			Input: &command.Scan{
				Table: &command.SimpleTable{
					Schema:  "mySchema",
					Table:   "myTable",
					Alias:   "myAlias",
					Indexed: true,
					Index:   "myIndex",
				},
			},
		},
	},
	{
		// PROJECT
		&Command_Project{
			Cols: []*Column{
				{
					Table: "myTable1",
					Column: &Expr{
						Expr: &Expr_Literal{
							&LiteralExpr{
								Value: "literal",
							},
						},
					},
					Alias: "myAlias1",
				},
				{
					Table: "myTable2",
					Column: &Expr{
						Expr: &Expr_Literal{
							&LiteralExpr{
								Value: "literal",
							},
						},
					},
					Alias: "myAlias2",
				},
			},
			Input: &List{
				List: &List_Scan{
					Scan: &Command_Scan{
						Table: &SimpleTable{
							Schema:  "mySchema",
							Table:   "myTable",
							Alias:   "myAlias",
							Indexed: true,
							Index:   "myIndex",
						},
					},
				},
			},
		},
		&command.Project{
			Cols: []command.Column{
				{
					Table: "myTable1",
					Column: &command.LiteralExpr{
						Value: "literal",
					},
					Alias: "myAlias1",
				},
				{
					Table: "myTable2",
					Column: &command.LiteralExpr{
						Value: "literal",
					},
					Alias: "myAlias2",
				},
			},
			Input: &command.Scan{
				Table: &command.SimpleTable{
					Schema:  "mySchema",
					Table:   "myTable",
					Alias:   "myAlias",
					Indexed: true,
					Index:   "myIndex",
				},
			},
		},
	},
	{
		// DELETE
		&Command_Delete{
			Table: &SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
			Filter: &Expr{
				Expr: &Expr_Binary{
					Binary: &BinaryExpr{
						Operator: "operator",
						Left: &Expr{
							Expr: &Expr_Literal{
								Literal: &LiteralExpr{
									Value: "leftLiteral",
								},
							},
						},
						Right: &Expr{
							Expr: &Expr_Literal{
								Literal: &LiteralExpr{
									Value: "rightLiteral",
								},
							},
						},
					},
				},
			},
		},
		command.Delete{
			Table: &command.SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
			Filter: &command.BinaryExpr{
				Operator: "operator",
				Left: &command.LiteralExpr{
					Value: "leftLiteral",
				},
				Right: &command.LiteralExpr{
					Value: "rightLiteral",
				},
			},
		},
	},
	{
		// DROP TABLE
		&CommandDrop{
			Target:   0,
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
		command.DropTable{
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
	},
	{
		// DROP VIEW
		&CommandDrop{
			Target:   1,
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
		command.DropView{
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
	},
	{
		// DROP INDEX
		&CommandDrop{
			Target:   2,
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
		command.DropIndex{
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
	},
	{
		// DROP TRIGGER
		&CommandDrop{
			Target:   3,
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
		command.DropTrigger{
			IfExists: true,
			Schema:   "mySchema",
			Name:     "tableName",
		},
	},
	{
		// UPDATE
		&Command_Update{
			UpdateOr: 0,
			Table: &SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
			Updates: []*UpdateSetter{
				{
					Cols: []string{
						"col1",
						"col2",
					},
					Value: &UpdateSetter_Constant{
						Constant: &ConstantBooleanExpr{
							Value: true,
						},
					},
				},
			},
			Filter: &Expr{
				Expr: &Expr_Equality{
					Equality: &EqualityExpr{
						Left: &Expr{
							Expr: &Expr_Literal{
								Literal: &LiteralExpr{
									Value: "leftLiteral",
								},
							},
						},
						Right: &Expr{
							Expr: &Expr_Literal{
								Literal: &LiteralExpr{
									Value: "rightLiteral",
								},
							},
						},
					},
				},
			},
		},
		command.Update{
			UpdateOr: 0,
			Table: &command.SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
			Updates: []command.UpdateSetter{
				{
					Cols: []string{
						"col1",
						"col2",
					},
					Value: command.ConstantBooleanExpr{
						Value: true,
					},
				},
			},
			Filter: &command.EqualityExpr{
				Left: &command.LiteralExpr{
					Value: "leftLiteral",
				},
				Right: &command.LiteralExpr{
					Value: "rightLiteral",
				},
			},
		},
	},
	{
		// JOIN
		&Command_Join{
			Natural: true,
			Type:    0,
			Filter: &Expr{
				Expr: &Expr_Func{
					Func: &FunctionExpr{
						Name:     "function",
						Distinct: true,
						Args: []*Expr{
							{
								Expr: &Expr_Range{
									&RangeExpr{
										Needle: &Expr{
											Expr: &Expr_Literal{
												Literal: &LiteralExpr{
													Value: "literal",
												},
											},
										},
										Lo: &Expr{
											Expr: &Expr_Literal{
												Literal: &LiteralExpr{
													Value: "literal",
												},
											},
										},
										Hi: &Expr{
											Expr: &Expr_Literal{
												Literal: &LiteralExpr{
													Value: "literal",
												},
											},
										},
										Invert: false,
									},
								},
							},
							{
								Expr: &Expr_Unary{
									Unary: &UnaryExpr{
										Operator: "operator",
										Value: &Expr{
											Expr: &Expr_Literal{
												Literal: &LiteralExpr{
													Value: "literal",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			Left: &List{
				List: &List_Scan{
					Scan: &Command_Scan{
						Table: &SimpleTable{
							Schema:  "mySchema",
							Table:   "myTable",
							Alias:   "myAlias",
							Indexed: true,
							Index:   "myIndex",
						},
					},
				},
			},
			Right: &List{
				List: &List_Scan{
					Scan: &Command_Scan{
						Table: &SimpleTable{
							Schema:  "mySchema",
							Table:   "myTable",
							Alias:   "myAlias",
							Indexed: true,
							Index:   "myIndex",
						},
					},
				},
			},
		},
		command.Join{
			Natural: true,
			Type:    0,
			Filter: &command.FunctionExpr{
				Name:     "function",
				Distinct: true,
				Args: []command.Expr{
					&command.RangeExpr{
						Needle: &command.LiteralExpr{
							Value: "literal",
						},
						Lo: &command.LiteralExpr{
							Value: "literal",
						},
						Hi: &command.LiteralExpr{
							Value: "literal",
						},
						Invert: false,
					},
					&command.UnaryExpr{
						Operator: "operator",
						Value: &command.LiteralExpr{
							Value: "literal",
						},
					},
				},
			},
			Left: &command.Scan{
				Table: &command.SimpleTable{
					Schema:  "mySchema",
					Table:   "myTable",
					Alias:   "myAlias",
					Indexed: true,
					Index:   "myIndex",
				},
			},
			Right: &command.Scan{
				Table: &command.SimpleTable{
					Schema:  "mySchema",
					Table:   "myTable",
					Alias:   "myAlias",
					Indexed: true,
					Index:   "myIndex",
				},
			},
		},
	},
	{
		// LIMIT
		&Command_Limit{
			Limit: &Expr{
				Expr: &Expr_Literal{
					Literal: &LiteralExpr{
						Value: "literal",
					},
				},
			},
			Input: &List{
				List: &List_Scan{
					Scan: &Command_Scan{
						Table: &SimpleTable{
							Schema:  "mySchema",
							Table:   "myTable",
							Alias:   "myAlias",
							Indexed: true,
							Index:   "myIndex",
						},
					},
				},
			},
		},
		command.Limit{
			Limit: &command.LiteralExpr{
				Value: "literal",
			},
			Input: &command.Scan{
				Table: &command.SimpleTable{
					Schema:  "mySchema",
					Table:   "myTable",
					Alias:   "myAlias",
					Indexed: true,
					Index:   "myIndex",
				},
			},
		},
	},
	{
		// INSERT
		&Command_Insert{
			InsertOr: 0,
			Table: &SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
			Cols: []*Column{
				{
					Table: "myTable1",
					Column: &Expr{
						Expr: &Expr_Literal{
							&LiteralExpr{
								Value: "literal",
							},
						},
					},
					Alias: "myAlias1",
				},
				{
					Table: "myTable2",
					Column: &Expr{
						Expr: &Expr_Literal{
							&LiteralExpr{
								Value: "literal",
							},
						},
					},
					Alias: "myAlias2",
				},
			},
			DefaultValues: false,
			Input: &List{
				List: &List_Scan{
					Scan: &Command_Scan{
						Table: &SimpleTable{
							Schema:  "mySchema",
							Table:   "myTable",
							Alias:   "myAlias",
							Indexed: true,
							Index:   "myIndex",
						},
					},
				},
			},
		},
		command.Insert{
			InsertOr: 0,
			Table: &command.SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
			Cols: []command.Column{
				{
					Table: "myTable1",
					Column: &command.LiteralExpr{
						Value: "literal",
					},
					Alias: "myAlias1",
				},
				{
					Table: "myTable2",
					Column: &command.LiteralExpr{
						Value: "literal",
					},
					Alias: "myAlias2",
				},
			},
			DefaultValues: false,
			Input: &command.Scan{
				Table: &command.SimpleTable{
					Schema:  "mySchema",
					Table:   "myTable",
					Alias:   "myAlias",
					Indexed: true,
					Index:   "myIndex",
				},
			},
		},
	},
}

func TestConvertMessageToCommand(t *testing.T) {
	for _, tt := range messageToCommandTests {
		t.Run(tt.in.Kind().String(), func(t *testing.T) {
			msg := ConvertMessageToCommand(tt.in)
			if !reflect.DeepEqual(msg, tt.out) {
				t.Errorf("got %q, want %q", msg, tt.out)
			}
		})
	}
}
