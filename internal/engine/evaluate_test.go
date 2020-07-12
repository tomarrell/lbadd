package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/engine/types"
)

func TestFullTableScan(t *testing.T) {
	t.Skip("full table scan not implemented yet")

	assert := assert.New(t)

	e := createEngineOnEmptyDatabase(t)
	result, err := e.Evaluate(command.Scan{
		Table: command.SimpleTable{
			Table: "myTable",
		},
	})
	assert.NoError(err)
	assert.Equal(Table{}, result)
}

func TestProjection(t *testing.T) {
	assert := assert.New(t)

	e := createEngineOnEmptyDatabase(t)
	result, err := e.Evaluate(command.Project{
		Cols: []command.Column{
			{
				Column: command.LiteralExpr{Value: "column2"},
			},
		},
		Input: command.Values{
			Values: [][]command.Expr{
				{command.LiteralExpr{Value: "hello"}, command.LiteralExpr{Value: "world"}, command.ConstantBooleanExpr{Value: true}},
				{command.LiteralExpr{Value: "foo"}, command.LiteralExpr{Value: "bar"}, command.ConstantBooleanExpr{Value: false}},
			},
		},
	})
	assert.NoError(err)
	assert.Equal(Table{
		Cols: []Col{
			{
				QualifiedName: "column2",
				Type:          types.String,
			},
		},
		Rows: []Row{
			{
				Values: []types.Value{types.NewString("world")},
			},
			{
				Values: []types.Value{types.NewString("bar")},
			},
		},
	}, result)
}
