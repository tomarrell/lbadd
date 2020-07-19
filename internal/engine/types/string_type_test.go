package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringType_Compare(t *testing.T) {
	tests := []struct {
		name    string
		first   Value
		second  Value
		want    int
		wantErr bool
	}{
		{
			"nils",
			nil,
			nil,
			0,
			true,
		},
		{
			"empty strings",
			NewString(""),
			NewString(""),
			0,
			false,
		},
		{
			"simple",
			NewString("a"),
			NewString("b"),
			-1,
			false,
		},
		{
			"equal",
			NewString("f"),
			NewString("f"),
			0,
			false,
		},
		{
			"long",
			NewString("foh382w9fh3wo4rgefisawel"),
			NewString("9548h7gor8shuspdofjwepor"),
			1,
			false,
		},
		{
			"different",
			NewString("abc"),
			NewString("z"),
			-1,
			false,
		},
		{
			"one empty",
			NewString("abc"),
			NewString(""),
			1,
			false,
		},
		{
			"uncomparable",
			NewDate(time.Now()),
			NewString(""),
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("ltr", func(t *testing.T) {
				assert := assert.New(t)
				res, err := String.Compare(tt.first, tt.second)
				if tt.wantErr {
					assert.Error(err)
				} else {
					assert.Equal(tt.want, res)
				}
			})
			t.Run("rtl", func(t *testing.T) {
				assert := assert.New(t)
				res, err := String.Compare(tt.second, tt.first)
				if tt.wantErr {
					assert.Error(err)
				} else {
					assert.Equal(tt.want, res*-1)
				}
			})
		})
	}
}

func TestStringType_Cast(t *testing.T) {
	tests := []struct {
		name    string
		from    Value
		want    Value
		wantErr bool
	}{
		{
			"string to string",
			NewString("abc"),
			NewString("abc"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			typ := tt.want.Type()
			got, err := typ.(Caster).Cast(tt.from)
			assert.Equal(tt.want, got)
			if tt.wantErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
		})
	}
}
