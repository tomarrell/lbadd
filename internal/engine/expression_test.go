package engine

import (
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/engine/types"
)

type evaluateExpressionTest struct {
	name    string
	e       Engine
	ctx     ExecutionContext
	expr    command.Expr
	want    types.Value
	wantErr string
}

func TestEngine_evaluateExpression(t *testing.T) {
	fixedTimestamp, err := time.Parse("2006-01-02T15:04:05", "2020-06-01T14:05:12")
	assert.NoError(t, err)
	fixedTimeProvider := func() time.Time { return fixedTimestamp }

	testEvaluateExpressionTest(t, []evaluateExpressionTest{
		{
			"nil",
			builder().build(),
			newEmptyExecutionContext(),
			nil,
			nil,
			"cannot evaluate expression of type <nil>",
		},
		{
			"true",
			builder().build(),
			newEmptyExecutionContext(),
			command.ConstantBooleanExpr{Value: true},
			types.NewBool(true),
			"",
		},
		{
			"false",
			builder().build(),
			newEmptyExecutionContext(),
			command.ConstantBooleanExpr{Value: false},
			types.NewBool(false),
			"",
		},
	})
	t.Run("functions", func(t *testing.T) {
		testEvaluateExpressionTest(t, []evaluateExpressionTest{
			{
				"function NOW",
				builder().
					timeProvider(fixedTimeProvider).
					build(),
				newEmptyExecutionContext(),
				command.FunctionExpr{
					Name: "NOW",
				},
				types.NewDate(fixedTimestamp),
				"",
			},
			{
				"unknown function",
				builder().build(),
				newEmptyExecutionContext(),
				command.FunctionExpr{
					Name: "NOTEXIST",
				},
				nil,
				"no function for name NOTEXIST(...)",
			}})
	})
	t.Run("arithmetic", func(t *testing.T) {
		t.Run("op=add", func(t *testing.T) {
			testEvaluateExpressionTest(t, []evaluateExpressionTest{
				{
					"simple addition",
					builder().build(),
					newEmptyExecutionContext(),
					command.BinaryExpr{
						Left:     command.LiteralExpr{Value: "5"},
						Operator: "+",
						Right:    command.LiteralExpr{Value: "6"},
					},
					types.NewInteger(11),
					"",
				},
			})
		})
		t.Run("op=sub", func(t *testing.T) {
			testEvaluateExpressionTest(t, []evaluateExpressionTest{
				{
					"simple subtraction",
					builder().build(),
					newEmptyExecutionContext(),
					command.BinaryExpr{
						Left:     command.LiteralExpr{Value: "6"},
						Operator: "-",
						Right:    command.LiteralExpr{Value: "5"},
					},
					types.NewInteger(1),
					"",
				},
			})
		})
		t.Run("op=mul", func(t *testing.T) {
			testEvaluateExpressionTest(t, []evaluateExpressionTest{
				{
					"simple multiplication",
					builder().build(),
					newEmptyExecutionContext(),
					command.BinaryExpr{
						Left:     command.LiteralExpr{Value: "6"},
						Operator: "*",
						Right:    command.LiteralExpr{Value: "5"},
					},
					types.NewInteger(30),
					"",
				},
			})
		})
		t.Run("op=div", func(t *testing.T) {
			testEvaluateExpressionTest(t, []evaluateExpressionTest{
				{
					"simple division",
					builder().build(),
					newEmptyExecutionContext(),
					command.BinaryExpr{
						Left:     command.LiteralExpr{Value: "15"},
						Operator: "/",
						Right:    command.LiteralExpr{Value: "5"},
					},
					types.NewReal(3),
					"",
				},
			})
		})
		t.Run("op=mod", func(t *testing.T) {
			testEvaluateExpressionTest(t, []evaluateExpressionTest{
				{
					"simple modulo",
					builder().build(),
					newEmptyExecutionContext(),
					command.BinaryExpr{
						Left:     command.LiteralExpr{Value: "7"},
						Operator: "%",
						Right:    command.LiteralExpr{Value: "5"},
					},
					types.NewInteger(2),
					"",
				},
			})
		})
		t.Run("op=pow", func(t *testing.T) {
			testEvaluateExpressionTest(t, []evaluateExpressionTest{
				{
					"simple exponentiation",
					builder().build(),
					newEmptyExecutionContext(),
					command.BinaryExpr{
						Left:     command.LiteralExpr{Value: "2"},
						Operator: "**",
						Right:    command.LiteralExpr{Value: "4"},
					},
					types.NewInteger(16),
					"",
				},
			})
		})
	})
}

func testEvaluateExpressionTest(t *testing.T, tests []evaluateExpressionTest) {
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
