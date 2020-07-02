package engine

import (
	"reflect"
	"testing"

	"github.com/tomarrell/lbadd/internal/engine/types"
)

func Test_builtinMax(t *testing.T) {
	type args struct {
		args []types.Value
	}
	tests := []struct {
		name    string
		args    args
		want    types.Value
		wantErr bool
	}{
		{
			"empty",
			args{
				[]types.Value{},
			},
			nil,
			false,
		},
		{
			"bools",
			args{
				[]types.Value{
					types.NewBool(true),
					types.NewBool(false),
					types.NewBool(false),
					types.NewBool(true),
					types.NewBool(false),
					types.NewBool(true),
					types.NewBool(false),
					types.NewBool(true),
					types.NewBool(true),
				},
			},
			types.NewBool(true),
			false,
		},
		{
			"integers",
			args{
				[]types.Value{
					types.NewInteger(3456),
					types.NewInteger(0),
					types.NewInteger(76526),
					types.NewInteger(1),
					types.NewInteger(23685245),
					types.NewInteger(45634),
					types.NewInteger(1345),
					types.NewInteger(346),
					types.NewInteger(5697),
				},
			},
			types.NewInteger(23685245),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinMax(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinMax() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("builtinMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
