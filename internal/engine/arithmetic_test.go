package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/engine/types"
)

func TestEngine_add(t *testing.T) {
	type args struct {
		ctx   ExecutionContext
		left  types.Value
		right types.Value
	}
	tests := []struct {
		name    string
		args    args
		want    types.Value
		wantErr string
	}{
		{
			"nils",
			args{
				newEmptyExecutionContext(),
				nil,
				nil,
			},
			nil,
			"cannot add <nil> and <nil>",
		},
		{
			"left nil",
			args{
				newEmptyExecutionContext(),
				nil,
				types.NewInteger(5),
			},
			nil,
			"cannot add <nil> and types.IntegerValue",
		},
		{
			"right nil",
			args{
				newEmptyExecutionContext(),
				types.NewInteger(5),
				nil,
			},
			nil,
			"cannot add types.IntegerValue and <nil>",
		},
		{
			"simple",
			args{
				newEmptyExecutionContext(),
				types.NewInteger(5),
				types.NewInteger(6),
			},
			types.NewInteger(11),
			"",
		},
		{
			"zero",
			args{
				newEmptyExecutionContext(),
				types.NewInteger(0),
				types.NewInteger(0),
			},
			types.NewInteger(0),
			"",
		},
		{
			"both negative",
			args{
				newEmptyExecutionContext(),
				types.NewInteger(-5),
				types.NewInteger(-5),
			},
			types.NewInteger(-10),
			"",
		},
		{
			"left negative",
			args{
				newEmptyExecutionContext(),
				types.NewInteger(-5),
				types.NewInteger(10),
			},
			types.NewInteger(5),
			"",
		},
		{
			"right negative",
			args{
				newEmptyExecutionContext(),
				types.NewInteger(10),
				types.NewInteger(-5),
			},
			types.NewInteger(5),
			"",
		},
		{
			"overflow",
			args{
				newEmptyExecutionContext(),
				types.NewInteger((1 << 63) - 1),
				types.NewInteger(5),
			},
			types.NewInteger(-(1 << 63) + 4),
			"",
		},
		{
			"negative overflow",
			args{
				newEmptyExecutionContext(),
				types.NewInteger(-(1 << 63)),
				types.NewInteger(-1),
			},
			types.NewInteger((1 << 63) - 1),
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			e := createEngineOnEmptyDatabase(t)
			got, err := e.add(tt.args.ctx, tt.args.left, tt.args.right)
			if tt.wantErr != "" {
				assert.EqualError(err, tt.wantErr)
			} else {
				assert.NoError(err)
			}
			assert.Equal(tt.want, got)
		})
	}
}
