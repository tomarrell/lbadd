package engine

import (
	"testing"

	"github.com/rs/zerolog"
)

func TestEngine_cmp(t *testing.T) {
	tests := []struct {
		name  string
		left  Value
		right Value
		want  cmpResult
	}{
		{
			"true <-> true",
			BoolValue{Value: true},
			BoolValue{Value: true},
			cmpEqual,
		},
		{
			"true <-> false",
			BoolValue{Value: true},
			BoolValue{Value: false},
			cmpGreaterThan,
		},
		{
			"false <-> true",
			BoolValue{Value: false},
			BoolValue{Value: true},
			cmpLessThan,
		},
		{
			"false <-> false",
			BoolValue{Value: false},
			BoolValue{Value: false},
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
