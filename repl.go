package main

import (
	"bufio"
	"bytes"
	"io"
)

type repl struct {
	executor *executor
}

func newRepl() *repl {
	return &repl{
		executor: newExecutor(),
	}
}

func (r *repl) readCommand(reader io.Reader) (instruction, error) {
	sc := bufio.NewScanner(reader)
	sc.Split(splitBySpace)
	sc.Scan()

	instr := instruction{}

	cmd := newCommand(sc.Text())
	switch cmd {
	case commandInsert:
	case commandSelect:
	case commandDelete:
	}

	return instr, nil
}

func splitBySpace(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexByte(data, ' '); i >= 0 {
		return i + 1, dropSpace(data[0:i]), nil
	}

	if atEOF {
		return len(data), dropSpace(data), nil
	}

	return 0, nil, nil
}

func dropSpace(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == ' ' {
		return data[0 : len(data)-1]
	}
	return data
}
