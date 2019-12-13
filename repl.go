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
		instr.params = tokens[1:]
	case commandSelect:
		instr.command = commandSelect
		instr.params = tokens[1:]
	case commandDelete:
		instr.command = commandDelete
		instr.params = tokens[1:]
	default:
		return instr, nil
	}

	return instr, nil
}
