package optimization

import (
	"reflect"
	"testing"

	"github.com/tomarrell/lbadd/internal/compiler/command"
)

func TestOptHalfJoin(t *testing.T) {
	tests := []struct {
		name  string
		cmd   command.Command
		want  command.Command
		want1 bool
	}{
		{
			"not applicable",
			command.Select{
				Input: command.Scan{
					Table: command.SimpleTable{Table: "foobar"},
				},
			},
			nil,
			false,
		},
		{
			"optimize right",
			command.Select{
				Input: command.Join{
					Left: command.Scan{
						Table: command.SimpleTable{Table: "foobar"},
					},
					Right: nil,
				},
			},
			command.Select{
				Input: command.Scan{
					Table: command.SimpleTable{Table: "foobar"},
				},
			},
			true,
		},
		{
			"optimize left",
			command.Select{
				Input: command.Join{
					Left: nil,
					Right: command.Scan{
						Table: command.SimpleTable{Table: "foobar"},
					},
				},
			},
			command.Select{
				Input: command.Scan{
					Table: command.SimpleTable{Table: "foobar"},
				},
			},
			true,
		},
		{
			"optimize left deep single",
			command.Select{
				Input: command.Join{
					Left: command.Join{
						Left: nil,
						Right: command.Scan{
							Table: command.SimpleTable{Table: "a"},
						},
					},
					Right: command.Scan{
						Table: command.SimpleTable{Table: "b"},
					},
				},
			},
			command.Select{
				Input: command.Join{
					Left: command.Scan{
						Table: command.SimpleTable{Table: "a"},
					},
					Right: command.Scan{
						Table: command.SimpleTable{Table: "b"},
					},
				},
			},
			true,
		},
		{
			"optimize left deep double",
			command.Select{
				Input: command.Join{
					Left: nil,
					Right: command.Join{
						Left: nil,
						Right: command.Scan{
							Table: command.SimpleTable{Table: "c"},
						},
					},
				},
			},
			command.Select{
				Input: command.Scan{
					Table: command.SimpleTable{Table: "c"},
				},
			},
			true,
		},
		{
			"nil join",
			command.Select{
				Input: command.Join{
					Left:  nil,
					Right: nil,
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := OptHalfJoin(tt.cmd)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OptHalfJoin() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("OptHalfJoin() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
