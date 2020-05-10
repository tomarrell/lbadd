package raft

import (
	"sync"

	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

// Server represents a raft server.
type Server interface {
	//NewServer returns a node variable initialised with all raft parameters.
	NewServer(conn network.Conn) (nodes *Node)
	// LeaderElection function starts a leader election from a single node in the cluster.
	// It returns an error based on what happened if it cannot start the election.
	// The function caller doesn't need to wait for a voting response from this function,
	// the function triggers the necessary functions responsible to continue the raft cluster
	// into it's working stage if the node won the election.
	LeaderElection(node *Node) error
	// RequestVoteResponse function is called on a request from a candidate for a vote. This function
	// generates the response for the responder node to send back to the candidate node.
	RequestVoteResponse(node *Node, req *message.RequestVoteRequest) *message.RequestVoteResponse
	// AppendEntriesResponse function is called on a request from the leader to append log data
	// to the follower node. This function generates the response to be sent to the leader node.
	AppendEntriesResponse(node *Node, req *message.AppendEntriesRequest) *message.AppendEntriesResponse
}

// Available states
const (
	LeaderState    = "leader"
	CandidateState = "candidate"
	FollowerState  = "follower"
)

// Node describes the current state of a raft node.
// The raft paper describes this as a "State" but node
// seemed more intuitive.
type Node struct {
	State string

	PersistentState     *PersistentState
	VolatileState       *VolatileState
	VolatileStateLeader *VolatileStateLeader

	// log zerolog.Logger
}

// PersistentState describes the persistent state data on a raft node.
type PersistentState struct {
	CurrentTerm int32
	VotedFor    []byte
	Log         []message.LogData

	SelfID  []byte
	SelfIP  network.Conn
	PeerIPs []network.Conn
	mu      sync.Mutex
}

// VolatileState describes the volatile state data on a raft node.
type VolatileState struct {
	CommitIndex int
	LastApplied int
}

// VolatileStateLeader describes the volatile state data that exists on a raft leader.
type VolatileStateLeader struct {
	NextIndex  []int // Holds the nextIndex value for each of the followers in the cluster.
	MatchIndex []int // Holds the matchIndex value for each of the followers in the cluster.
}

// NewRaftCluster initialises a raft cluster with the given configuration.
func NewRaftCluster(cluster cluster.Cluster) []*Node {
	var clusterNodes []*Node

	for i := range cluster.Nodes() {
		node := &Node{
			PersistentState: &PersistentState{
				CurrentTerm: 0,
				VotedFor:    nil,
				SelfIP:      cluster.Nodes()[i],
				PeerIPs:     cluster.Nodes(),
			},
			VolatileState: &VolatileState{
				CommitIndex: -1,
				LastApplied: -1,
			},
			VolatileStateLeader: &VolatileStateLeader{},
		}

		clusterNodes = append(clusterNodes, node)
	}
	return clusterNodes
}
