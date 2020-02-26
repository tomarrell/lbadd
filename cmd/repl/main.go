package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/tomarrell/lbadd/internal/cli"
	"github.com/tomarrell/lbadd/internal/executor"
	"github.com/tomarrell/lbadd/internal/server"
)

const (
	// ExitAbnormal is the exit code that the application will return upon
	// internal abnormal exit.
	ExitAbnormal = 1
	// ExitInterrupt is the exit code that the application will return when
	// aborted by the user.
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
		verbose   = flags.Bool("verbose", false, "enable verbose output")
		logfile   = flags.String("logfile", "lbadd.cli.log", "define a log file to write messages to")
		port      = flags.Int("port", 34213, "publish the database server on this port")
		host      = flags.String("host", "", "publish the database server on this host")
		headless  = flags.Bool("headless", false, "don't use a cli")
		noConsole = flags.Bool("quiet", false, "don't print logs to stdout")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return fmt.Errorf("parse flags: %w", err)
	}

	// open the log file
	file, err := os.OpenFile(*logfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		return fmt.Errorf("open logfile: %w", err)
	}
	defer file.Close()

	// programCtx is the context, that all components should run on. When
	// invoking cancel, all started components should stop processing.
	programCtx, cancel := context.WithCancel(context.Background())

	// initialize all writers
	writers := []io.Writer{
		// performant file writer
		diode.NewWriter(
			file, // output writer
			1000, // pool size
			0,    // poll interval, use a waiter
			func(missed int) {
				_, _ = fmt.Fprintf(stderr, "Logger is falling behind, skipping %d messages\n", missed)
			},
		),
	}
	if !*noConsole {
		writers = append(writers,
			// unperformant console writer
			zerolog.ConsoleWriter{
				Out: stdout,
			},
		)
	}

	// initialize the root logger
	log := zerolog.New(io.MultiWriter(writers...)).With().
		Timestamp().
		Logger().
		Level(zerolog.InfoLevel)

	log.Info().
		Bool("headless", *headless).
		Msg("start new application session")

	// apply the verbose flag
	if *verbose {
		log = log.Level(zerolog.TraceLevel)
	}

	// print all flags on debug level
	log.Debug().
		Bool("verbose", *verbose).
		Str("logfile", *logfile).
		Int("publish", *port).
		Bool("headless", *headless).
		Bool("quiet", *noConsole).
		Msg("settings")

	// start listening for signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	go func() {
		select {
		case sig := <-signalChan: // first signal, cancel context
			log.Info().
				Str("signal", sig.String()).
				Msg("Attempting graceful shutdown, press again to force quit")
			cancel()
		case <-programCtx.Done():
		}
		<-signalChan // second signal, hard exit
		log.Info().
			Msg("Forcing shutdown")
		os.Exit(ExitInterrupt)
	}()

	// setup the executor
	executor := executor.New(log.With().Str("component", "executor").Logger())

	// setup server endpoint
	server := server.New(log.With().Str("component", "server").Logger(), executor)
	runServer := func() {
		if err := server.ListenAndServe(programCtx, *host+":"+strconv.Itoa(int(*port))); err != nil {
			log.Error().
				Err(err).
				Msg("listen and serve failed")
		}
	}

	/*
		Handle headless.

		If headless, start and wait for the server, otherwise start the server
		in the background and start the cli. stdin will not be used if
		headless=true.
	*/
	if *headless {
		runServer()
	} else {
		go runServer()

		// run the cli
		cli := cli.New(programCtx, stdin, stdout, executor)
		cli.Start()
	}

	return nil
}
