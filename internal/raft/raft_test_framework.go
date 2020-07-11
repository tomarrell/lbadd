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
	log         zerolog.Logger
	parameters  OperationParameters
	config      NetworkConfiguration
	opChannel   chan OpData
	execChannel chan OpData
	roundsChan  chan bool
	opQueue     []OpData
	round       int
	shutdown    chan bool
}

// NewSimpleRaftTest provides a ready to use raft test framework.
func NewSimpleRaftTest(
	log zerolog.Logger,
	parameters OperationParameters,
	config NetworkConfiguration) *SimpleRaftTest {
	opChan := make(chan OpData, 5)
	execChan := make(chan OpData, 5)
	shutdownChan := make(chan bool, 1)
	roundsChan := make(chan bool, 1)
	return &SimpleRaftTest{
		log:         log,
		parameters:  parameters,
		config:      config,
		opChannel:   opChan,
		execChannel: execChan,
		roundsChan:  roundsChan,
		opQueue:     []OpData{},
		round:       0,
		shutdown:    shutdownChan,
	}
}

// OpParams returns the parameters of operations of the test.
func (t *SimpleRaftTest) OpParams() OperationParameters {
	return t.parameters
}

// Config returns the configuration under which the test is running.
func (t *SimpleRaftTest) Config() NetworkConfiguration {
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

	// start the execution goroutine.
	log.Debug().Msg("beginning execution goroutine")
	go t.executeOperation()

	// Look for incoming operations and parallely run them
	// while waiting for the limit of the execution.
	// Once the limit of the execution is reached, wait for
	// all operations to finish and end the test.
	for {
		select {
		case data := <-t.opChannel:
			log.Debug().
				Str("executing", string(data.Op)).
				Msg("beginning execution")
			go t.execute(data)
		case <-shutDownTimer.C:
			return t.GracefulShutdown()
		case <-t.roundsChan:
			return t.GracefulShutdown()
		}
	}
}

// GracefulShutdown shuts down all operations of the server after waiting
// all running operations to complete while not accepting any more op reqs.
func (t *SimpleRaftTest) GracefulShutdown() error {
	t.shutdown <- true
	log.Debug().
		Msg("gracefully shutting down")
	return nil
}

// InjectOperation initiates an operation in the raft cluster based on the args.
func (t *SimpleRaftTest) InjectOperation(op Operation, args interface{}) {
	// check whether test has begun.

	opData := OpData{
		Op:   op,
		Data: args,
	}
	log.Debug().Msg("injecting operation")
	t.opChannel <- opData
}

// execute appends the operation to the queue which will
// be cleared in definite intervals.
func (t *SimpleRaftTest) execute(opData OpData) {
	log.Debug().Msg("operation moved to execution channel")
	t.execChannel <- opData
}

// executeOperation is always ready to run an incoming operation.
// It looks for the shutdown signal from the hook channel and
// shutdown by not allowing further operations to execute.
//
// When both cases of the select statement recieve a signal,
// select chooses one at random. This doesn't affect the operation
// as the execution will shutdown right after that operation is
// completed.
func (t *SimpleRaftTest) executeOperation() {
	for {
		select {
		case <-t.shutdown:
			log.Debug().Msg("execution shutting down")
			return
		case operation := <-t.execChannel:
			log.Debug().Msg("executing operation")
			switch operation.Op {
			case SendData:
				d := operation.Data.(*OpSendData)
				t.SendData(d)
			case StopNode:
				d := operation.Data.(*OpStopNode)
				t.StopNode(d)
			case PartitionNetwork:
				d := operation.Data.(*OpPartitionNetwork)
				t.PartitionNetwork(d)
			case RestartNode:
				d := operation.Data.(*OpRestartNode)
				t.RestartNode(d)
			}
		default:
			continue
		}
	}
}

// func (t *SimpleRaftTest) roundHook() {
// 	t.round++
// 	t.roundsChan <- true
// }

// SendData sends command data to the cluster by calling
// the appropriate function in the raft module.
func (t *SimpleRaftTest) SendData(d *OpSendData) {

}

// StopNode stops the given node in the network.
// This is a test of robustness in the system to recover from
// a failure of a node.
//
// The implementation can involve killing/stopping the
// respective node.
func (t *SimpleRaftTest) StopNode(d *OpStopNode) {

}

// PartitionNetwork partitions the network into one or more
// groups as dictated by the arguments. This means that the
// nodes in different groups cannot communicate with the
// nodes in a different group.
//
// The implementation can involve removing the nodes in the
// in the respective "cluster" variable so that they are no
// longer available to access it.
func (t *SimpleRaftTest) PartitionNetwork(d *OpPartitionNetwork) {

}

// RestartNode restarts a previously stopped node which has
// all resources allocated to it but went down for any reason.
func (t *SimpleRaftTest) RestartNode(d *OpRestartNode) {

}
