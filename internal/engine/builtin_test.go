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
					types.BoolValue{Value: true},
					types.BoolValue{Value: false},
					types.BoolValue{Value: false},
					types.BoolValue{Value: true},
					types.BoolValue{Value: false},
					types.BoolValue{Value: true},
					types.BoolValue{Value: false},
					types.BoolValue{Value: true},
					types.BoolValue{Value: true},
				},
			},
			types.BoolValue{Value: true},
			false,
		},
		{
			"numbers",
			args{
				[]types.Value{
					types.NumericValue{Value: 32698.7236},
					types.NumericValue{Value: 33020.4705},
					types.NumericValue{Value: 28550.057},
					types.NumericValue{Value: 17980.2620},
					types.NumericValue{Value: 37105.784},
					types.NumericValue{Value: 623164325426457348.4231},
					types.NumericValue{Value: 53226.854},
					types.NumericValue{Value: 49344.266},
					types.NumericValue{Value: 10634.3037},
					types.NumericValue{Value: 36735.083},
					types.NumericValue{Value: 14828.1674},
				},
			},
			types.NumericValue{Value: 623164325426457348.4231},
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
