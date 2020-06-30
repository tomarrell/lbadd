package message

import (
	"testing"

	"github.com/tomarrell/lbadd/internal/compiler/command"
)

// var commandToMessageTests = []struct {
// 	in  command.Command
// 	out Message
// }{
// 	{
// 		command.Scan{
// 			Table: command.SimpleTable{
// 				Schema:  "mySchema",
// 				Table:   "myTable",
// 				Alias:   "myAlias",
// 				Indexed: true,
// 				Index:   "myIndex",
// 			},
// 		},
// 		&Command_Scan{
// 			Table: &SimpleTable{
// 				Schema:  "mySchema",
// 				Table:   "myTable",
// 				Alias:   "myAlias",
// 				Indexed: true,
// 				Index:   "myIndex",
// 			},
// 		},
// 	},
// }

// func Test_CommandToMessage(t *testing.T) {
// 	t.SkipNow()
// 	for _, tt := range commandToMessageTests {
// 		t.Run(tt.in.String(), func(t *testing.T) {
// 			msg := ConvertCommandToMessage(tt.in)
// 			if msg != tt.out {
// 				t.Errorf("got %q, want %q", msg, tt.out)
// 			}
// 		})
// 	}
// }

var messageToCommandTests = []struct {
	in  Message
	out command.Command
}{
	{
		&Command_Scan{
			Table: &SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
		},
		command.Scan{
			Table: command.SimpleTable{
				Schema:  "mySchema",
				Table:   "myTable",
				Alias:   "myAlias",
				Indexed: true,
				Index:   "myIndex",
			},
		},
	},
	// {
	// 	&Command_Select{
	// 		Filter: &Expr{
	// 			Expr: ,
	// 		},
	// 	},
	// 	command.Select{},
	// },
}

func Test_MessageToCommand(t *testing.T) {
	t.SkipNow()
	for _, tt := range messageToCommandTests {
		t.Run(tt.in.Kind().String(), func(t *testing.T) {
			msg := ConvertMessageToCommand(tt.in)
			if msg != tt.out {
				t.Errorf("got %q, want %q", msg, tt.out)
			}
		})
	}
}
