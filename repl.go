package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type repl struct {
	executor *executor
}

func newRepl() *repl {
	return &repl{
		executor: newExecutor(),
	}
}

func (r *repl) start() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Starting Bad SQL repl")

	for {
		fmt.Print("$ ")
		sc.Scan()

		instr, err := r.readCommand(sc.Text())
		if err != nil {
			fmt.Printf("\nInvalid command: %v", err)
			continue
		}

		r.executor.execute(instr)
	}
}

func (r *repl) readCommand(input string) (instruction, error) {
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
