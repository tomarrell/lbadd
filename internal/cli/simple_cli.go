package cli

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"github.com/tomarrell/lbadd/internal/executor"
	"github.com/tomarrell/lbadd/internal/executor/command"
	"github.com/tomarrell/lbadd/internal/parser"
)

var _ Cli = (*simpleCli)(nil)

type simpleCli struct {
	ctx    context.Context
	closed bool

	in  io.Reader
	out io.Writer

	scanner *bufio.Scanner

	exec executor.Executor
}

func newSimpleCli(ctx context.Context, in io.Reader, out io.Writer, exec executor.Executor) *simpleCli {
	return &simpleCli{
		ctx:     ctx,
		in:      in,
		out:     out,
		scanner: bufio.NewScanner(in),
		exec:    exec,
	}
}

func (c *simpleCli) Start() {
	lines := make(chan string)

	// read from the cli scanner
	go func() {
		for c.scanner.Scan() {
			lines <- c.scanner.Text()
		}
	}()

	for !c.closed {
		_, _ = fmt.Fprint(c.out, "$ ")
		select {
		case line := <-lines:
			c.handleInput(line)
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *simpleCli) handleInput(input string) {
	switch input {
	case "help", "h", "?", "\\?":
		c.printHelp()
	case "q", "exit", "\\q":
		c.quit()
	default:
		c.handleSQL(input)
	}
}

func (c *simpleCli) printHelp() {
	_, _ = fmt.Fprintln(c.out, "Available Commands:\n// TODO")
}

func (c *simpleCli) quit() {
	_, _ = fmt.Fprintln(c.out, "Bye!")
	c.closed = true
}

func (c *simpleCli) handleSQL(sqlInput string) {
	parser := parser.New(sqlInput)
	for c.ctx.Err() == nil {
		// parse the input statement
		stmt, errs, ok := parser.Next()
		if !ok {
			break
		}
		// print errors to the output
		for _, err := range errs {
			_, _ = fmt.Fprintf(c.out, "error while parsing command: %v\n", err)
		}
		// if there were errors, abandon the statement
		if len(errs) > 0 {
			_, _ = fmt.Fprintf(c.out, "will skip statement, because there were %d parse errors\n", len(errs))
			continue
		}
		// convert AST to IR
		command, err := command.From(stmt)
		if err != nil {
			_, _ = fmt.Fprintf(c.out, "error while compiling command: %v\n", err)
			continue
		}
		// execute the command
		result, err := c.exec.Execute(command)
		if err != nil {
			_, _ = fmt.Fprintf(c.out, "error while executing command: %v\n", err)
			continue
		}
		// print the result of the command execution to the output
		_, _ = fmt.Fprintln(c.out, result)
	}
}
