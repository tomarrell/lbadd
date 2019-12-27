package lbadd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_executor_execute(t *testing.T) {
	type fields struct {
		db *db
	}
	type args struct {
		instr instruction
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    result
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "create table creates a new empty table",
			fields: fields{db: &db{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &executor{
				db: tt.fields.db,
			}

			got, err := e.execute(tt.args.instr)
			if err != nil {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
