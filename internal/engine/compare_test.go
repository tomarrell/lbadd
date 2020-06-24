package engine

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/engine/types"
)

func TestEngine_cmp(t *testing.T) {
	tests := []struct {
		name  string
		left  types.Value
		right types.Value
		want  cmpResult
	}{
		{
			"true <-> true",
			types.BoolValue{Value: true},
			types.BoolValue{Value: true},
			cmpEqual,
		},
		{
			"true <-> false",
			types.BoolValue{Value: true},
			types.BoolValue{Value: false},
			cmpGreaterThan,
		},
		{
			"false <-> true",
			types.BoolValue{Value: false},
			types.BoolValue{Value: true},
			cmpLessThan,
		},
		{
			"false <-> false",
			types.BoolValue{Value: false},
			types.BoolValue{Value: false},
			cmpEqual,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Engine{
				log: zerolog.Nop(),
			}
			if got := e.cmp(tt.left, tt.right); got != tt.want {
				t.Errorf("Engine.cmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
