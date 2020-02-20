package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/tomarrell/lbadd/internal/cli"
	"github.com/tomarrell/lbadd/internal/executor"
)

const (
	ExitAbnormal  = 1
	ExitInterrupt = 2
)

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(ExitAbnormal)
	}
}

func run(args []string, stdin io.Reader, stdout, stderr io.Writer) error {
	// setup flags
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var (
		verbose = flags.Bool("verbose", false, "enable verbose output")
	)
	_ = *verbose // TODO: use *verbose to configure a logger

	if err := flags.Parse(args[1:]); err != nil {
		return fmt.Errorf("parse flags: %w", err)
	}

	// programCtx is the context, that all components should run on. When
	// invoking cancel, all started components should stop processing.
	programCtx, cancel := context.WithCancel(context.Background())

	// start listening for signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
			_, _ = fmt.Fprintln(stdout, "Attempting graceful shutdown, press again to force quit")
		case <-programCtx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(ExitInterrupt)
	}()

	// run the cli
	cli := cli.New(programCtx, stdin, stdout, executor.New())
	cli.Start()

	return nil
}
