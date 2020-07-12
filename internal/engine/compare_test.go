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
			types.NewBool(true),
			types.NewBool(true),
			cmpEqual,
		},
		{
			"true <-> false",
			types.NewBool(true),
			types.NewBool(false),
			cmpGreaterThan,
		},
		{
			"false <-> true",
			types.NewBool(false),
			types.NewBool(true),
			cmpLessThan,
		},
		{
			"false <-> false",
			types.NewBool(false),
			types.NewBool(false),
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
