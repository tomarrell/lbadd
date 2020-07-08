package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolType_Compare(t *testing.T) {
	type args struct {
		left  Value
		right Value
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr string
	}{
		{
			"nils",
			args{nil, nil},
			0,
			"type mismatch: want Bool, got <nil>",
		},
		{
			"null <-> nil",
			args{NewNull(Bool), nil},
			0,
			"type mismatch: want Bool, got <nil>",
		},
		{
			"null <-> null",
			args{NewNull(Bool), NewNull(Bool)},
			-1,
			"",
		},
		{
			"null <-> true",
			args{NewNull(Bool), NewBool(true)},
			-1,
			"",
		},
		{
			"null <-> false",
			args{NewNull(Bool), NewBool(false)},
			-1,
			"",
		},
		{
			"left nil",
			args{nil, NewBool(false)},
			0,
			"type mismatch: want Bool, got <nil>",
		},
		{
			"right nil",
			args{NewBool(false), nil},
			0,
			"type mismatch: want Bool, got <nil>",
		},
		{
			"equal true",
			args{NewBool(true), NewBool(true)},
			0,
			"",
		},
		{
			"equal false",
			args{NewBool(false), NewBool(false)},
			0,
			"",
		},
		{
			"less",
			args{NewBool(false), NewBool(true)},
			-1,
			"",
		},
		{
			"greater",
			args{NewBool(true), NewBool(false)},
			1,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			got, err := Bool.Compare(tt.args.left, tt.args.right)
			if tt.wantErr != "" {
				assert.EqualError(err, tt.wantErr)
			} else {
				assert.NoError(err)
			}
			assert.Equal(tt.want, got)
		})
	}
}

func TestBoolType_Deserialize(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    Value
		wantErr string
	}{
		{
			"nil",
			nil,
			nil,
			"unexpected data size 0, need 1",
		},
		{
			"too large",
			[]byte{1, 2},
			nil,
			"unexpected data size 2, need 1",
		},
		{
			"too large",
			[]byte{1, 2, 3},
			nil,
			"unexpected data size 3, need 1",
		},
		{
			"true",
			[]byte{1},
			NewBool(true),
			"",
		},
		{
			"true",
			[]byte{2},
			NewBool(true),
			"",
		},
		{
			"true",
			[]byte{^uint8(0)},
			NewBool(true),
			"",
		},
		{
			"false",
			[]byte{0},
			NewBool(false),
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			got, err := Bool.Deserialize(tt.data)
			if tt.wantErr != "" {
				assert.EqualError(err, tt.wantErr)
			} else {
				assert.NoError(err)
			}
			assert.Equal(tt.want, got)
		})
	}
}

func TestBoolType_Serialize(t *testing.T) {
	tests := []struct {
		name    string
		v       Value
		want    []byte
		wantErr string
	}{
		{
			"nil",
			nil,
			nil,
			"type mismatch: want Bool, got <nil>",
		},
		{
			"string",
			NewString("foobar"),
			nil,
			"type mismatch: want Bool, got String",
		},
		{
			"true",
			NewBool(true),
			[]byte{0x01},
			"",
		},
		{
			"false",
			NewBool(false),
			[]byte{0x00},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			got, err := Bool.Serialize(tt.v)
			if tt.wantErr != "" {
				assert.EqualError(err, tt.wantErr)
			} else {
				assert.NoError(err)
			}
			assert.Equal(tt.want, got)
		})
	}
}
