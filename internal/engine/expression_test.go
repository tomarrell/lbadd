package engine

import (
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/engine/types"
)

func TestEngine_evaluateExpression(t *testing.T) {
	fixedTimestamp, err := time.Parse("2006-01-02T15:04:05", "2020-06-01T14:05:12")
	assert.NoError(t, err)
	fixedTimeProvider := func() time.Time { return fixedTimestamp }

	tests := []struct {
		name    string
		e       Engine
		ctx     ExecutionContext
		expr    command.Expr
		want    types.Value
		wantErr string
	}{
		{
			"nil",
			builder().build(),
			ExecutionContext{},
			nil,
			nil,
			"cannot evaluate expression of type <nil>",
		},
		{
			"true",
			builder().build(),
			ExecutionContext{},
			command.ConstantBooleanExpr{Value: true},
			types.NewBool(true),
			"",
		},
		{
			"false",
			builder().build(),
			ExecutionContext{},
			command.ConstantBooleanExpr{Value: false},
			types.NewBool(false),
			"",
		},
		{
			"function NOW",
			builder().
				timeProvider(fixedTimeProvider).
				build(),
			ExecutionContext{},
			command.FunctionExpr{
				Name: "NOW",
			},
			types.NewDate(fixedTimestamp),
			"",
		},
		{
			"unknown function",
			builder().build(),
			ExecutionContext{},
			command.FunctionExpr{
				Name: "NOTEXIST",
			},
			nil,
			"no function for name NOTEXIST(...)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			got, err := tt.e.evaluateExpression(tt.ctx, tt.expr)
			assert.Equal(tt.want, got)
			if tt.wantErr != "" {
				assert.EqualError(err, tt.wantErr)
			} else {
				assert.NoError(err)
			}
		})
	}
}

type engineBuilder struct {
	e Engine
}

func builder() engineBuilder {
	return engineBuilder{
		Engine{
			log: zerolog.Nop(),
		},
	}
}

func (b engineBuilder) timeProvider(tp timeProvider) engineBuilder {
	b.e.timeProvider = tp
	return b
}

func (b engineBuilder) build() Engine {
	return b.e
}
