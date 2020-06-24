package engine

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/engine/storage"
)

func TestEngine(t *testing.T) {
	assert := assert.New(t)

	fs := afero.NewMemMapFs()
	f, err := fs.Create("mydbfile")
	assert.NoError(err)
	dbFile, err := storage.Create(f)
	assert.NoError(err)

	e, err := New(dbFile)
	assert.NoError(err)
	result, err := e.Evaluate(command.Values{
		Values: [][]command.Expr{
			{command.LiteralExpr{Value: "hello"}, command.LiteralExpr{Value: "world"}, command.ConstantBooleanExpr{Value: true}},
			{command.LiteralExpr{Value: "foo"}, command.LiteralExpr{Value: "bar"}, command.ConstantBooleanExpr{Value: false}},
		},
	})
	assert.NoError(err)
	assert.NotNil(result)

	// check cols
	cols := result.Cols()
	assert.Len(cols, 3)
	// col[0]
	assert.Equal(2, cols[0].Size())
	assert.Equal(stringType, cols[0].Type())
	// col[1]
	assert.Equal(2, cols[1].Size())
	assert.Equal(stringType, cols[1].Type())
	// col[2]
	assert.Equal(2, cols[2].Size())
	assert.Equal(numericType, cols[2].Type())
	// col value types
	assert.Equal(cols[0].Type(), cols[0].Get(0).Type())
	assert.Equal(cols[0].Type(), cols[0].Get(1).Type())
	assert.Equal(cols[1].Type(), cols[1].Get(0).Type())
	assert.Equal(cols[1].Type(), cols[1].Get(1).Type())
	assert.Equal(cols[2].Type(), cols[2].Get(0).Type())
	assert.Equal(cols[2].Type(), cols[2].Get(1).Type())

	// check rows
	rows := result.Rows()
	assert.Len(rows, 2)
	// row[0]
	assert.Equal(3, rows[0].Size())
	assert.Equal("hello", rows[0].Get(0).(StringValue).Value)
	assert.Equal("world", rows[0].Get(1).(StringValue).Value)
	assert.Equal(true, rows[0].Get(2).(BoolValue).Value)
	// row[0]
	assert.Equal(3, rows[1].Size())
	assert.Equal("foo", rows[1].Get(0).(StringValue).Value)
	assert.Equal("bar", rows[1].Get(1).(StringValue).Value)
	assert.Equal(false, rows[1].Get(2).(BoolValue).Value)
}
