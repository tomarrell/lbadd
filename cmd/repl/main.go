package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/tomarrell/lbadd"
)

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdin io.Reader, stdout, stderr io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var (
		verbose = flags.Bool("verbose", false, "enable verbose output")
	)
	_ = *verbose // TODO: use *verbose to configure a logger

	if err := flags.Parse(args[1:]); err != nil {
		return fmt.Errorf("parse flags: %w", err)
	}

	r := lbadd.NewRepl()
	r.Start()

	return nil
}
