package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/spf13/cobra"
	"github.com/tomarrell/lbadd/internal/executor"
	"github.com/tomarrell/lbadd/internal/master"
	"github.com/tomarrell/lbadd/internal/worker"
)

// intended to be set in build process
var (
	Version = "master"
)

// application constants
const (
	ApplicationName = "lbadd"
)

// exit codes
const (
	// ExitAbnormal is the exit code that the application will return upon
	// internal abnormal exit.
	ExitAbnormal = 1
	// ExitInterrupt is the exit code that the application will return when
	// aborted by the user.
	ExitInterrupt = 2
)

type ctxKey uint8

const (
	ctxKeyStdin ctxKey = iota
	ctxKeyStdout
	ctxKeyStderr
	ctxKeyLog
)

// command line arguments
var (
	verbose bool
	logfile string
	addr    string
)

// documentation strings
const (
	rootCmdShortDoc = "Start the database application"
	rootCmdLongDoc  = ""

	versionCmdShortDoc = "Print version information about this executable"
	versionCmdLongDoc  = ""

	startCmdShortDoc = "Start either a master or a worker node"
	startCmdLongDoc  = ""

	startMasterCmdShortDoc = "Start a master node"
	startMasterCmdLongDoc  = `Start a master node on the address that is specified in the addr flag.
This will start an lbadd master node on the specified address,
waiting for incoming connections from lbadd worker nodes.`

	startWorkerCmdShortDoc = "Start a worker node"
	startWorkerCmdLongDoc  = ""
)

var (
	rootCmd = &cobra.Command{
		Use:   ApplicationName,
		Short: rootCmdShortDoc,
		Long:  rootCmdLongDoc,
		Args:  cobra.NoArgs,
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: versionCmdShortDoc,
		Long:  versionCmdLongDoc,
		Run:   printVersion,
		Args:  cobra.NoArgs,
	}

	startCmd = &cobra.Command{
		Use:   "start",
		Short: startCmdShortDoc,
		Long:  startCmdLongDoc,
		Args:  cobra.NoArgs,
	}

	startMasterCmd = &cobra.Command{
		Use:   "master",
		Short: startMasterCmdShortDoc,
		Long:  startMasterCmdLongDoc,
		Run:   startMaster,
		Args:  cobra.NoArgs,
	}

	startWorkerCmd = &cobra.Command{
		Use:   "worker",
		Short: startWorkerCmdShortDoc,
		Long:  startWorkerCmdLongDoc,
		Run:   startWorker,
		Args:  cobra.NoArgs,
	}
)

func init() {
	rootCmd.AddCommand(startCmd, versionCmd)
	startCmd.AddCommand(startMasterCmd, startWorkerCmd)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "print more logs")

	startCmd.PersistentFlags().StringVar(&logfile, "logfile", "lbadd.log", "define a log file to write logs to")

	startMasterCmd.PersistentFlags().StringVar(&addr, "addr", ":34213", "serve the database on this address")

	startWorkerCmd.PersistentFlags().StringVar(&addr, "addr", ":34213", "connect to a master node on this address")
}

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(ExitAbnormal)
	}
}

func run(args []string, stdin io.Reader, stdout, stderr io.Writer) error {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	log := createLogger(stdin, stdout, stderr)
	ctx = context.WithValue(ctx, ctxKeyLog, log)

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
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		log.Info().
			Msg("Forcing shutdown")
		os.Exit(ExitInterrupt)
	}()

	ctx = context.WithValue(ctx, ctxKeyStdin, stdin)
	ctx = context.WithValue(ctx, ctxKeyStdout, stdout)
	ctx = context.WithValue(ctx, ctxKeyStderr, stderr)

	return rootCmd.ExecuteContext(ctx)
}

func printVersion(cmd *cobra.Command, args []string) {
	stdout := cmd.Context().Value(ctxKeyStdout).(io.Writer)
	_, _ = fmt.Fprintf(stdout, "%s version %s\n", ApplicationName, Version)
}

func startMaster(cmd *cobra.Command, args []string) {
	log := cmd.Context().Value(ctxKeyLog).(zerolog.Logger)

	exec := createExecutor(log)

	masterLog := log.With().
		Str("component", "master").
		Logger()

	masterNode := master.New(masterLog, exec)
	if err := masterNode.ListenAndServe(cmd.Context(), addr); err != nil {
		log.Error().
			Err(err).
			Msg("listen and serve")
		os.Exit(ExitAbnormal)
	}
}

func startWorker(cmd *cobra.Command, args []string) {
	log := cmd.Context().Value(ctxKeyLog).(zerolog.Logger)

	exec := createExecutor(log)

	workerLog := log.With().
		Str("component", "worker").
		Logger()

	workerNode := worker.New(workerLog, exec)
	if err := workerNode.Connect(cmd.Context(), addr); err != nil {
		log.Error().
			Err(err).
			Msg("connect")
		os.Exit(ExitAbnormal)
	}
}

func createLogger(stdin io.Reader, stdout, stderr io.Writer) zerolog.Logger {
	// open the log file
	file, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		_, _ = fmt.Fprintln(stderr, fmt.Errorf("open logfile: %w", err).Error())
		os.Exit(ExitAbnormal)
	}

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
		// unperformant console writer
		zerolog.ConsoleWriter{
			Out: stdout,
		},
	}

	// initialize the root logger
	log := zerolog.New(io.MultiWriter(writers...)).With().
		Timestamp().
		Logger().
		Level(zerolog.InfoLevel)

	// apply the verbose flag
	if verbose {
		log = log.Level(zerolog.TraceLevel)
	}

	return log
}

func createExecutor(log zerolog.Logger) executor.Executor {
	execLog := log.With().
		Str("component", "executor").
		Logger()

	exec := executor.New(execLog)
	return exec
}
