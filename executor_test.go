package lbadd

import "testing"

func Test_executor_execute(t *testing.T) {
	type (
		fields struct {
			db *db
		}
		args struct {
			instr instruction
		}
	)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &executor{db: tt.fields.db}
			if err := e.execute(tt.args.instr); (err != nil) != tt.wantErr {
				t.Errorf("executor.execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
