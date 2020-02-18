package cli

import (
	"bufio"
	"fmt"
	"io"

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

		c.handleCommand(c.scanner.Text())

		_, _ = fmt.Fprintln(c.out, "")
	}
}

func (c *simpleCli) handleCommand(command string) {
	switch command {
	case "help", "h", "?", "\\?":
		fmt.Print("Available Commands:\n// TODO")
		return
	case "q", "exit", "\\q":
		fmt.Print("Bye!")
		return
	}

	parser := parser.New(command)
	for {
		stmt, errs, ok := parser.Next()
		if !ok {
			break
		}
		for _, err := range errs {
			_, _ = fmt.Fprintf(c.out, "error while parsing command: %v\n", err)
		}
		// TODO: define intermediary representation, convert and then execute
		if err := c.exec.Execute(stmt); err != nil {
			_, _ = fmt.Fprintf(c.out, "error while executing command: %v\n", err)
		}
	}
}

func (c *simpleCli) Close() error {
	c.closed = true
	return nil
}
