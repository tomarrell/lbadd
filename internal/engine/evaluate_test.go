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

func TestEngine_evaluateSelection(t *testing.T) {
	tests := []struct {
		name    string
		ctx     ExecutionContext
		sel     command.Select
		want    Table
		wantErr string
	}{
		{
			"trivial",
			newEmptyExecutionContext(),
			command.Select{
				Filter: command.ConstantBooleanExpr{Value: true},
				Input: command.Values{
					Values: [][]command.Expr{
						{command.LiteralExpr{Value: "hello"}, command.LiteralExpr{Value: "5"}, command.ConstantBooleanExpr{Value: true}},
						{command.LiteralExpr{Value: "foo"}, command.LiteralExpr{Value: "7"}, command.ConstantBooleanExpr{Value: false}},
					},
				},
			},
			Table{
				Cols: []Col{
					{
						QualifiedName: "column1",
						Type:          types.String,
					},
					{
						QualifiedName: "column2",
						Type:          types.Integer,
					},
					{
						QualifiedName: "column3",
						Type:          types.Bool,
					},
				},
				Rows: []Row{
					{
						Values: []types.Value{types.NewString("hello"), types.NewInteger(5), types.NewBool(true)},
					},
					{
						Values: []types.Value{types.NewString("foo"), types.NewInteger(7), types.NewBool(false)},
					},
				},
			},
			"",
		},
		{
			"simple",
			newEmptyExecutionContext(),
			command.Select{
				Filter: command.EqualityExpr{
					Left:  command.LiteralExpr{Value: "column2"},
					Right: command.LiteralExpr{Value: "7"},
				},
				Input: command.Values{
					Values: [][]command.Expr{
						{command.LiteralExpr{Value: "hello"}, command.LiteralExpr{Value: "5"}, command.ConstantBooleanExpr{Value: true}},
						{command.LiteralExpr{Value: "foo"}, command.LiteralExpr{Value: "7"}, command.ConstantBooleanExpr{Value: false}},
					},
				},
			},
			Table{
				Cols: []Col{
					{
						QualifiedName: "column1",
						Type:          types.String,
					},
					{
						QualifiedName: "column2",
						Type:          types.Integer,
					},
					{
						QualifiedName: "column3",
						Type:          types.Bool,
					},
				},
				Rows: []Row{
					{
						Values: []types.Value{types.NewString("foo"), types.NewInteger(7), types.NewBool(false)},
					},
				},
			},
			"",
		},
		{
			"erronous filter",
			newEmptyExecutionContext(),
			command.Select{
				Filter: command.LiteralExpr{Value: "erronous"},
				Input: command.Values{
					Values: [][]command.Expr{
						{command.LiteralExpr{Value: "hello"}, command.LiteralExpr{Value: "5"}, command.ConstantBooleanExpr{Value: true}},
						{command.LiteralExpr{Value: "foo"}, command.LiteralExpr{Value: "7"}, command.ConstantBooleanExpr{Value: false}},
					},
				},
			},
			Table{},
			"cannot use command.LiteralExpr as filter",
		},
		{
			"column against column",
			newEmptyExecutionContext(),
			command.Select{
				Filter: command.EqualityExpr{
					Left:  command.LiteralExpr{Value: "column2"},
					Right: command.LiteralExpr{Value: "column1"},
				},
				Input: command.Values{
					Values: [][]command.Expr{
						{command.LiteralExpr{Value: "hello"}, command.LiteralExpr{Value: "world"}},
						{command.LiteralExpr{Value: "foo"}, command.LiteralExpr{Value: "foo"}},
					},
				},
			},
			Table{
				Cols: []Col{
					{
						QualifiedName: "column1",
						Type:          types.String,
					},
					{
						QualifiedName: "column2",
						Type:          types.String,
					},
				},
				Rows: []Row{
					{
						Values: []types.Value{types.NewString("foo"), types.NewString("foo")},
					},
				},
			},
			"",
		},
		{
			"column against string",
			newEmptyExecutionContext(),
			command.Select{
				Filter: command.EqualityExpr{
					Left:  command.LiteralExpr{Value: "column2"},
					Right: command.LiteralExpr{Value: "world"},
				},
				Input: command.Values{
					Values: [][]command.Expr{
						{command.LiteralExpr{Value: "hello"}, command.LiteralExpr{Value: "world"}},
						{command.LiteralExpr{Value: "foo"}, command.LiteralExpr{Value: "foo"}},
					},
				},
			},
			Table{
				Cols: []Col{
					{
						QualifiedName: "column1",
						Type:          types.String,
					},
					{
						QualifiedName: "column2",
						Type:          types.String,
					},
				},
				Rows: []Row{
					{
						Values: []types.Value{types.NewString("hello"), types.NewString("world")},
					},
				},
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			e := createEngineOnEmptyDatabase(t)
			got, err := e.evaluateSelection(tt.ctx, tt.sel)
			if tt.wantErr != "" {
				assert.EqualError(err, tt.wantErr)
			} else {
				assert.NoError(err)
			}
			assert.Equal(tt.want, got)
		})
	}
}
