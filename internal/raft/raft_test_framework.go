package raft

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Usage of framework:
//
// * Create a new instance of the RaftTestFramework;
//   this begins the raft cluster and all operations.
// * Call monitor to start seeing real time logs.
// * Use "InjectData" to insert a log into the cluster.
// * Use "InjectOperation" with appropriate args to
//   trigger an operation in the cluster.1
var _ TestFramework = (*SimpleRaftTest)(nil)

// SimpleRaftTest implements TestFramework.
type SimpleRaftTest struct {
	log        zerolog.Logger
	parameters OperationParameters
	config     NetworkConfiguration
	opChannel  chan OpData
}

// NewSimpleRaftTest provides a ready to use raft test framework.
func NewSimpleRaftTest(
	log zerolog.Logger,
	parameters OperationParameters,
	config NetworkConfiguration) *SimpleRaftTest {
	opChan := make(chan OpData)
	return &SimpleRaftTest{
		log,
		parameters,
		config,
		opChan,
	}
}

// OpParams returns the parameters of operations of the test.
func (t *SimpleRaftTest) OpParams() OperationParameters {
	return t.parameters
}

// Cfg returns the configuration under which the test is running.
func (t *SimpleRaftTest) Cfg() NetworkConfiguration {
	return t.config
}

// BeginTest starts all the cluster operations by creating and
// starting the cluster and the nodes. This operation will be
// completely stable and allows failing of servers underneath
// while monitoring their behavior.
//
// BeginTest will wrapped under a Go Test for ease of use.
func (t *SimpleRaftTest) BeginTest() error {
	// Check for proper config before beginning.
	// Start the cluster and blah.
	//
	//
	shutDownTimer := time.NewTimer(time.Duration(t.OpParams().TimeLimit) * time.Second)
	signal := make(chan bool, 1)

	go func() {
		<-shutDownTimer.C
		signal <- true
	}()

	go func() {
		for {
			// If current rounds == required rounds
			// signal <- true
			return
		}
	}()

	for {
		select {
		case <-t.opChannel:
		case <-signal:
			return t.GracefulShutdown()
		}
	}
}

// GracefulShutdown shuts down all operations of the server after waiting
// all running operations to complete while not accepting any more op reqs.
func (t *SimpleRaftTest) GracefulShutdown() error {
	log.Debug().
		Msg("gracefully shutting down")

	return nil
}

// InjectOperation initiates an operation in the raft cluster based on the args.
func (t *SimpleRaftTest) InjectOperation(op Operation, args interface{}) {
	switch op {
	case SendData:
	case StopNode:
	case PartitionNetwork:
	}
}

//Monitor monitors
func (t *SimpleRaftTest) Monitor() error {
	return nil
}
