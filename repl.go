package lbadd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Repl is an interactive print loop which accepts instructions in the form of
// the database's intermediary representation, and executes the statements
// against the database.
type Repl struct {
	executor *executor
}

// NewRepl creates a new repl instance
func NewRepl() *Repl {
	return &Repl{
		executor: newExecutor(exeConfig{
			order: defaultOrder,
		}),
	}
}

// Start begings the execution of the given repl instance
func (r *Repl) Start() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Starting Bad SQL repl")

	for {
		fmt.Print("$ ")
		sc.Scan()

		input := sc.Text()
		switch input {
		case "help", "h", "?", "\\?":
			fmt.Println(`Available Commands:
// TODO`)
		case "q", "exit", "\\q":
			fmt.Println("Bye!")
			return
		}

		instr, err := r.readCommand(input)
		if err != nil {
			fmt.Printf("\nInvalid command: %v", err)
			continue
		}

		_, err = r.executor.execute(instr)
		if err != nil {
			fmt.Printf("Err: %v\n", err)
			continue
		}
	}
}

func (r *Repl) readCommand(input string) (instruction, error) {
	tokens := strings.Split(input, " ")
	instr := instruction{}

	switch newCommand(tokens[0]) {
	case commandInsert:
		instr.command = commandInsert
		instr.table = tokens[1]
		instr.params = tokens[2:]
	case commandSelect:
		instr.command = commandSelect
		instr.table = tokens[1]
		instr.params = tokens[2:]
	case commandDelete:
		instr.command = commandDelete
		instr.table = tokens[1]
		instr.params = tokens[2:]
	default:
		return instr, nil
	}

	return instr, nil
}
