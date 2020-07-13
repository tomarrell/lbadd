package engine

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/engine/storage"
	"github.com/tomarrell/lbadd/internal/engine/types"
)

func TestEngine(t *testing.T) {
	assert := assert.New(t)

	e := createEngineOnEmptyDatabase(t)

	result, err := e.Evaluate(command.Values{
		Values: [][]command.Expr{
			{command.LiteralExpr{Value: "hello"}, command.LiteralExpr{Value: "world"}, command.ConstantBooleanExpr{Value: true}},
			{command.LiteralExpr{Value: "foo"}, command.LiteralExpr{Value: "bar"}, command.ConstantBooleanExpr{Value: false}},
		},
	})
	assert.NoError(err)
	assert.NotNil(result)

	// check rows
	rows := result.Rows
	assert.Len(rows, 2)
	// row[0]
	assert.Equal(3, len(rows[0].Values))
	assert.Equal("hello", rows[0].Values[0].(types.StringValue).Value)
	assert.Equal("world", rows[0].Values[1].(types.StringValue).Value)
	assert.Equal(true, rows[0].Values[2].(types.BoolValue).Value)
	// row[1]
	assert.Equal(3, len(rows[1].Values))
	assert.Equal("foo", rows[1].Values[0].(types.StringValue).Value)
	assert.Equal("bar", rows[1].Values[1].(types.StringValue).Value)
	assert.Equal(false, rows[1].Values[2].(types.BoolValue).Value)
}

func createEngineOnEmptyDatabase(t *testing.T) Engine {
	assert := assert.New(t)

	fs := afero.NewMemMapFs()
	f, err := fs.Create("mydbfile")
	assert.NoError(err)
	dbFile, err := storage.Create(f)
	assert.NoError(err)

	e, err := New(dbFile)
	assert.NoError(err)
	return e
}
