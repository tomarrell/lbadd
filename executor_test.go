package lbadd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/btree"
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
		// TODO add test cases
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

func Test_executor_executeCreateTable(t *testing.T) {
	type fields struct {
		db  *db
		cfg exeConfig
	}
	type args struct {
		instr instruction
	}

	order := 3

	tests := []struct {
		name       string
		fields     fields
		args       args
		want       result
		wantTables map[string]table
		wantErr    bool
	}{
		{
			name:    "creates a new empty table",
			fields:  fields{db: &db{tables: map[string]table{}}, cfg: exeConfig{order: order}},
			args:    args{instr: instruction{command: commandCreateTable, table: "users"}},
			want:    result{created: 1},
			wantErr: false,
			wantTables: map[string]table{
				"users": {
					name:    "users",
					store:   btree.NewBtreeOrder(order),
					columns: []column{},
				},
			},
		},
		{
			name:   "creates a new table with single column",
			fields: fields{db: &db{tables: map[string]table{}}, cfg: exeConfig{order: order}},
			args: args{instr: instruction{
				command: commandCreateTable,
				table:   "users",
				params:  []string{"name", "string", "false"},
			}},
			want:    result{created: 1},
			wantErr: false,
			wantTables: map[string]table{
				"users": {
					name:  "users",
					store: btree.NewBtreeOrder(order),
					columns: []column{
						{
							dataType:   columnTypeString,
							name:       "name",
							isNullable: false,
						},
					},
				},
			},
		},
		{
			name:   "creates a new table with multiple columns",
			fields: fields{db: &db{tables: map[string]table{}}, cfg: exeConfig{order: order}},
			args: args{instr: instruction{
				command: commandCreateTable,
				table:   "users",
				params:  []string{"name", "string", "false", "age", "integer", "true"},
			}},
			want:    result{created: 1},
			wantErr: false,
			wantTables: map[string]table{
				"users": {
					name:  "users",
					store: btree.NewBtreeOrder(order),
					columns: []column{
						{
							dataType:   columnTypeString,
							name:       "name",
							isNullable: false,
						},
						{
							dataType:   columnTypeInt,
							name:       "age",
							isNullable: true,
						},
					},
				},
			},
		},
		{
			name:   "fails to create if datatype is unknown",
			fields: fields{db: &db{tables: map[string]table{}}, cfg: exeConfig{order: order}},
			args: args{instr: instruction{
				command: commandCreateTable,
				table:   "users",
				params:  []string{"name", "unknown", "false"},
			}},
			want:       result{created: 0},
			wantErr:    true,
			wantTables: map[string]table{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &executor{
				db:  tt.fields.db,
				cfg: tt.fields.cfg,
			}

			got, err := e.executeCreateTable(tt.args.instr)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantTables, tt.fields.db.tables)
		})
	}
}

func Test_executor_executeSelect(t *testing.T) {
	type fields struct {
		db  *db
		cfg exeConfig
	}
	type args struct {
		instr instruction
	}

	order := 3
	mockTable := table{
		name:  "user",
		store: btree.NewBtreeOrder(order),
		columns: []column{
			{
				dataType:   columnTypeInt,
				name:       "id",
				isNullable: false,
			},
			{
				dataType:   columnTypeString,
				name:       "name",
				isNullable: false,
			},
		},
	}
	mockTable.store.Insert(0, "John Smith")

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    result
		wantErr bool
	}{
		{
			name: "error when table does not exist",
			fields: fields{
				db: &db{
					tables: map[string]table{mockTable.name: mockTable},
				},
				cfg: exeConfig{order: order},
			},
			args: args{
				instr: instruction{
					command: commandSelect,
					table:   "missing_table",
					params:  []string{},
				},
			},
			want:    result{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &executor{
				db:  tt.fields.db,
				cfg: tt.fields.cfg,
			}

			got, err := e.executeSelect(tt.args.instr)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
