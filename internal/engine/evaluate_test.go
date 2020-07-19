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

func TestEngine_evaluateProjection(t *testing.T) {
	tests := []struct {
		name    string
		ctx     ExecutionContext
		proj    command.Project
		want    Table
		wantErr string
	}{
		{
			"empty",
			newEmptyExecutionContext(),
			command.Project{
				Cols: []command.Column{},
				Input: command.Values{
					Values: [][]command.Expr{
						{command.LiteralExpr{Value: "hello"}, command.LiteralExpr{Value: "world"}, command.ConstantBooleanExpr{Value: true}},
						{command.LiteralExpr{Value: "foo"}, command.LiteralExpr{Value: "bar"}, command.ConstantBooleanExpr{Value: false}},
					},
				},
			},
			Table{
				Cols: []Col{},
				Rows: []Row{},
			},
			"",
		},
		{
			"simple",
			newEmptyExecutionContext(),
			command.Project{
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
			},
			Table{
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
			},
			"",
		},
		{
			"simple with alias",
			newEmptyExecutionContext(),
			command.Project{
				Cols: []command.Column{
					{
						Column: command.LiteralExpr{Value: "column2"},
						Alias:  "foo",
					},
				},
				Input: command.Values{
					Values: [][]command.Expr{
						{command.LiteralExpr{Value: "hello"}, command.LiteralExpr{Value: "world"}, command.ConstantBooleanExpr{Value: true}},
						{command.LiteralExpr{Value: "foo"}, command.LiteralExpr{Value: "bar"}, command.ConstantBooleanExpr{Value: false}},
					},
				},
			},
			Table{
				Cols: []Col{
					{
						QualifiedName: "column2",
						Alias:         "foo",
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
			},
			"",
		},
		{
			"missing column",
			newEmptyExecutionContext(),
			command.Project{
				Cols: []command.Column{
					{
						Column: command.LiteralExpr{Value: "foo"},
					},
				},
				Input: command.Values{
					Values: [][]command.Expr{
						{command.LiteralExpr{Value: "hello"}},
						{command.LiteralExpr{Value: "foo"}},
					},
				},
			},
			Table{},
			"no column with name or alias 'foo'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			e := createEngineOnEmptyDatabase(t)
			got, err := e.evaluateProjection(tt.ctx, tt.proj)
			if tt.wantErr != "" {
				assert.EqualError(err, tt.wantErr)
			} else {
				assert.NoError(err)
			}
			assert.Equal(tt.want, got)
		})
	}
}
