package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
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
		logfile = flags.String("logfile", "lbadd.cli.log", "define a log file to write messages to")
	)

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

	// Initialize a root logger
	file, err := os.OpenFile(*logfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		return fmt.Errorf("open logfile: %w", err)
	}

	log := zerolog.New(
		diode.NewWriter(
			file,                // output writer
			1000,                // pool size
			10*time.Millisecond, // poll interval
			func(missed int) {
				_, _ = fmt.Fprintf(stderr, "Logger is falling behind, skipping %d messages\n", missed)
			},
		),
	).With().
		Timestamp().
		Logger().
		Level(zerolog.InfoLevel)

	log.Info().Msg("start new session")

	// apply the verbose flag
	if *verbose {
		log = log.Level(zerolog.TraceLevel)
	}

	// run the cli
	cli := cli.New(programCtx, stdin, stdout, executor.New(log.With().Str("component", "executor").Logger()))
	cli.Start()

	return nil
}
