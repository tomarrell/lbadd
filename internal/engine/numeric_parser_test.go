package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/engine/types"
)

func TestToNumericValue(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		want  types.Value
		want1 bool
	}{
		{
			"empty",
			"",
			nil,
			false,
		},
		{
			"half hex",
			"0x",
			nil,
			false,
		},
		{
			"hex",
			"0x123ABC",
			types.NewInteger(0x123ABC),
			true,
		},
		{
			"hex",
			"0xFF",
			types.NewInteger(0xFF),
			true,
		},
		{
			"full hex spectrum",
			"0x0123456789ABCDEF",
			types.NewInteger(0x0123456789ABCDEF),
			true,
		},
		{
			"full hex spectrum",
			"0xFEDCBA987654321",
			types.NewInteger(0xFEDCBA987654321),
			true,
		},
		{
			"small integral",
			"0",
			types.NewInteger(0),
			true,
		},
		{
			"small integral",
			"1",
			types.NewInteger(1),
			true,
		},
		{
			"integral",
			"1234567",
			types.NewInteger(1234567),
			true,
		},
		{
			"integral",
			"42",
			types.NewInteger(42),
			true,
		},
		{
			"real",
			"0.0",
			types.NewReal(0),
			true,
		},
		{
			"real",
			".0",
			types.NewReal(0),
			true,
		},
		{
			"only decimal point",
			".",
			nil,
			false,
		},
		{
			"real with exponent",
			".0E2",
			types.NewReal(0),
			true,
		},
		{
			"real with exponent",
			"5.7E-242",
			types.NewReal(5.7E-242),
			true,
		},
		{
			"invalid exponent",
			".0e2", // lowercase 'e' is not allowed
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			got, ok := ToNumericValue(tt.s)
			assert.Equal(tt.want, got)
			assert.Equal(tt.want1, ok)
		})
	}
}
