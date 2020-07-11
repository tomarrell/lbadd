package raft

import (
	"github.com/tomarrell/lbadd/internal/compile"
	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
)

// TestFramework describes a testing framework for creating
// a complete integration testing of the different modules of a
// raft implementation.
//
// The framework allows injecting mutiple types of operations, like,
// stopping a node, partitioning the cluster, adding an entry to the log
// of the cluster whilst allowing to monitor the operations of the cluster.
//
// The framework will be completely stable and run obeying the Operation
// parameters while running the cluster operations and logging their behaviour.
type TestFramework interface {
	// OpParams provides the parameters for operation of the raft cluster.
	OpParams() OperationParameters
	// Cfg provides the network configuration of the cluster.
	Config() NetworkConfiguration
	// BeginTest kicks of all operations by starting the raft cluster.
	// It obeys the parameters of operation and raises an error if the
	// conditions for the test don't satisfy.
	BeginTest() error
	// InjectOperation will initiate the given operation in the cluster.
	InjectOperation(op Operation, args interface{})
	// GracefulShutdown ensures the cluster is shutdown by waiting for
	// all the running operations to complete.
	GracefulShutdown() error
}

// OperationParameters are the bounds which dictate the parameters
// for the running integration test.
//
// The raft operation will run until it reaches the first of the two bounds.
// If the operation is in consensus in the first round, the TimeLimit variable
// will stop the operation once it's reached.
//
// All operations will be stopped gracefully whenever possible after the bounds
// are reached.
type OperationParameters struct {
	// Rounds specifies how many rounds the raft operation must proceed until.
	Rounds int
	// TimeLimit specifies the limit until which the raft operation will run.
	TimeLimit int
}

// NetworkConfiguration holds the details of the network of the cluster.
type NetworkConfiguration struct {
	IPs []network.Conn
	IDs []id.ID
}

// Operation describes the different types of operations that can be performed on the cluster.
type Operation int

// Types of Operations.
const (
	SendData Operation = 1 + iota
	StopNode
	PartitionNetwork
	RestartNode
)

// OpData fully describes a runnable operation on the raft cluster.
// The "data" field can be either of the data related to the operation.
type OpData struct {
	Op   Operation
	Data interface{}
}

// OpSendData describes the data related to SendData.
type OpSendData struct {
	Data []*compile.Command
}

// OpStopNode describes the data related to StopNode.
type OpStopNode struct {
	NodeID id.ID
}

// OpPartitionNetwork describes the data related to PartitionNetwork.
type OpPartitionNetwork struct {
	Groups [][]id.ID
}

// OpRestartNode describes the data related to RestartNode.
type OpRestartNode struct {
	NodeID id.ID
}
