package engine

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/engine/types"
)

func TestEngine_evaluateExpression(t *testing.T) {
	tests := []struct {
		name    string
		expr    command.Expr
		want    types.Value
		wantErr bool
	}{
		{
			"nil",
			nil,
			nil,
			true,
		},
		{
			"true",
			command.ConstantBooleanExpr{Value: true},
			types.BoolValue{Value: true},
			false,
		},
		{
			"false",
			command.ConstantBooleanExpr{Value: false},
			types.BoolValue{Value: false},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			e := Engine{
				log: zerolog.Nop(),
			}
			got, err := e.evaluateExpression(tt.expr)
			assert.Equal(tt.want, got)
			if tt.wantErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
		})
	}
}
