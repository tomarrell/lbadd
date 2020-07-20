package raft

import (
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
)

// Usage of framework:
//
// * Create a new instance of the RaftTestFramework;
//   this begins the raft cluster and all operations.
//
//	 raftTest := NewSimpleRaftTest(log,opParams,cfg)
//
// * Use "InjectOperation" with appropriate args to
//   trigger an operation in the cluster.
var _ TestFramework = (*SimpleRaftTest)(nil)

// SimpleRaftTest implements TestFramework.
type SimpleRaftTest struct {
	log        zerolog.Logger
	parameters OperationParameters
	config     NetworkConfiguration
	raftNodes  []*SimpleServer

	opChannel   chan OpData
	execChannel chan OpData
	roundsChan  chan bool
	opQueue     []OpData
	round       int
	shutdown    chan bool
	mu          sync.Mutex
	cancelFunc  context.CancelFunc
}

// NewSimpleRaftTest provides a ready to use raft test framework.
func NewSimpleRaftTest(
	log zerolog.Logger,
	parameters OperationParameters,
	config NetworkConfiguration,
	raftNodes []*SimpleServer,
	cancel context.CancelFunc,
) *SimpleRaftTest {
	opChan := make(chan OpData, len(parameters.Operations))
	execChan := make(chan OpData, len(parameters.Operations))
	shutdownChan := make(chan bool, 4)
	roundsChan := make(chan bool, 4)
	return &SimpleRaftTest{
		log:         log,
		parameters:  parameters,
		config:      config,
		raftNodes:   raftNodes,
		opChannel:   opChan,
		execChannel: execChan,
		roundsChan:  roundsChan,
		opQueue:     []OpData{},
		round:       0,
		shutdown:    shutdownChan,
		cancelFunc:  cancel,
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
func (t *SimpleRaftTest) BeginTest(ctx context.Context) error {
	// if t.config.IDs == nil && t.config.IPs == nil {
	// 	return errors.New("nil network configuration")
	// }

	// start up the raft operation.
	for i := range t.raftNodes {
		go func(i int) {
			t.raftNodes[i].OnCompleteOneRound(t.roundHook)
			_ = t.raftNodes[i].Start(ctx)
		}(i)
	}

	shutDownTimer := time.NewTimer(time.Duration(t.OpParams().TimeLimit) * time.Second)

	// start the execution goroutine.
	log.Debug().Msg("beginning execution goroutine")
	go t.executeOperation()

	log.Debug().Msg("initiating operation injection")
	go t.pushOperations()

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
			log.Debug().
				Msg("shutting down - reached time limit")
			return t.GracefulShutdown()
		case <-t.roundsChan:
			log.Debug().
				Msg("shutting down - reached round limit")
			return t.GracefulShutdown()
		}
	}
}

// GracefulShutdown shuts down all operations of the server after waiting
// all running operations to complete while not accepting any more op reqs.
func (t *SimpleRaftTest) GracefulShutdown() error {
	t.cancelFunc()
	var errSlice multiError
	var errLock sync.Mutex
	for i := range t.raftNodes {
		err := t.raftNodes[i].Close()
		if err != nil {
			errLock.Lock()
			errSlice = append(errSlice, err)
			errLock.Unlock()
		}
	}
	if len(errSlice) != 0 {
		return errSlice
	}

	t.shutdown <- true
	log.Debug().
		Msg("gracefully shutting down")
	return nil
}

// InjectOperation initiates an operation in the raft cluster based on the args.
func (t *SimpleRaftTest) InjectOperation(op Operation, args interface{}) {
	opData := OpData{
		Op:   op,
		Data: args,
	}
	log.Debug().Msg("injecting operation")
	t.opChannel <- opData
}

// pushOperations pushes operations into the execution queue.
func (t *SimpleRaftTest) pushOperations() {
	for i := range t.parameters.Operations {
		t.opChannel <- t.parameters.Operations[i]
	}
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
		}
	}
}

func (t *SimpleRaftTest) roundHook() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.round++
	if t.round == t.parameters.Rounds*len(t.raftNodes) {
		t.roundsChan <- true
	}
}

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

func createRaftNodes(log zerolog.Logger, cluster *cluster.TestNetwork) []*SimpleServer {
	raftNodes := []*SimpleServer{}
	for i := range cluster.Clusters {
		node := NewServer(log, cluster.Clusters[i].Cluster)
		raftNodes = append(raftNodes, node)
	}

	return raftNodes
}
