package cli

import (
	"bufio"
	"fmt"
	"io"

	"github.com/tomarrell/lbadd/internal/executor/command"
	"github.com/tomarrell/lbadd/internal/parser"
)

var _ Cli = (*simpleCli)(nil)

type simpleCli struct {
	in  io.Reader
	out io.Writer

	closed  bool
	scanner *bufio.Scanner

	exec Executor
}

func newSimpleCli(in io.Reader, out io.Writer, exec Executor) *simpleCli {
	return &simpleCli{
		in:      in,
		out:     out,
		scanner: bufio.NewScanner(in),
		exec:    exec,
	}
}

func (c *simpleCli) Start() {
	for !c.closed {
		_, _ = fmt.Fprint(c.out, "$ ")
		if !c.scanner.Scan() {
			break
		}

		c.handleInput(c.scanner.Text())

		_, _ = fmt.Fprintln(c.out, "")
	}
}

func (c *simpleCli) handleInput(input string) {
	switch input {
	case "help", "h", "?", "\\?":
		fmt.Print("Available Commands:\n// TODO")
		return
	case "q", "exit", "\\q":
		fmt.Print("Bye!")
		return
	}

	parser := parser.New(input)
	for {
		stmt, errs, ok := parser.Next()
		if !ok {
			break
		}
		for _, err := range errs {
			_, _ = fmt.Fprintf(c.out, "error while parsing command: %v\n", err)
		}

		command, err := command.From(stmt)
		if err != nil {
			_, _ = fmt.Fprintf(c.out, "error while compiling command: %v\n", err)
		}
		if err := c.exec.Execute(command); err != nil {
			_, _ = fmt.Fprintf(c.out, "error while executing command: %v\n", err)
		}
	}
}

func (c *simpleCli) Close() error {
	c.closed = true
	return nil
}
