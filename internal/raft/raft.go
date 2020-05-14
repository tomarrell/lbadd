package raft

import (
	"context"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

// Available states
const (
	LeaderState    = "leader"
	CandidateState = "candidate"
	FollowerState  = "follower"
)

// Server is a description of a raft server.
type Server interface {
	Start() error
	OnReplication(ReplicationHandler)
	Input(string)
	io.Closer
}

// Cluster is a description of a cluster of servers.
type Cluster interface {
	OwnID() id.ID
	Nodes() []network.Conn
	Receive(context.Context) (network.Conn, message.Message, error)
	Broadcast(context.Context, message.Message) error
}

// ReplicationHandler is a handler setter.
type ReplicationHandler func(string)

// Node describes the current state of a raft node.
// The raft paper describes this as a "State" but node
// seemed more intuitive.
type Node struct {
	State      string
	LogChannel chan (message.LogData) // LogChannel is used to store the incoming logs from clients.

	PersistentState     *PersistentState
	VolatileState       *VolatileState
	VolatileStateLeader *VolatileStateLeader
}

// PersistentState describes the persistent state data on a raft node.
type PersistentState struct {
	CurrentTerm int32
	VotedFor    id.ID // VotedFor is nil at init, -1 if the node voted for itself and any number in the slice Nodes() to point at its voter.
	Log         []message.LogData

	SelfID   id.ID
	LeaderID id.ID          // LeaderID is nil at init, -1 if its the leader and any number in the slice Nodes() to point at the leader.
	PeerIPs  []network.Conn // PeerIPs has the connection variables of all the other nodes in the cluster.
	mu       sync.Mutex
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

var _ Server = (*simpleServer)(nil)

// simpleServer implements a server in a cluster.
type simpleServer struct {
	cluster       Cluster
	onReplication ReplicationHandler
	log           zerolog.Logger
}

// incomingData describes every request that the server gets.
type incomingData struct {
	conn network.Conn
	msg  message.Message
	err  error
}

// NewServer enables starting a raft server/cluster.
func NewServer(log zerolog.Logger, cluster Cluster) Server {
	return &simpleServer{
		log:     log.With().Str("component", "raft").Logger(),
		cluster: cluster,
	}
}

// Start starts a single raft node into beginning raft operations.
// This function starts the leader election and keeps a check on whether
// regular heartbeats to the node exists. It restarts leader election on failure to do so.
// This function also continuously listens on all the connections to the nodes
// and routes the requests to appropriate functions.
func (s *simpleServer) Start() (err error) {
	// Initialise all raft variables in this node.
	node := NewRaftCluster(s.cluster)
	ctx := context.Background()
	liveChan := make(chan *incomingData)
	// Listen forever on all node connections. This block of code checks what kind of
	// request has to be serviced and calls the necessary function to complete it.
	go func() {
		for {
			// Parallely start waiting for incoming data.
			go func() {
				conn, msg, err := s.cluster.Receive(ctx)
				liveChan <- &incomingData{
					conn,
					msg,
					err,
				}
			}()

			// If any sort of request (heartbeat,appendEntries,requestVote)
			// isn't received by the server(node) it restarts leader election.
			select {
			case <-randomTicker().C:
				StartElection(node)
			case data := <-liveChan:
				err := processIncomingData(data, node)
				if err != nil {
					return
				}
			}
		}
	}()

	// TODO: Just to maintain a blocking function.
	<-time.NewTicker(10000000 * time.Second).C
	return nil
}

func (s *simpleServer) OnReplication(handler ReplicationHandler) {
	s.onReplication = handler
}

func (s *simpleServer) Input(string) {

}

func (s *simpleServer) Close() error {
	return nil
}

// NewRaftCluster initialises a raft cluster with the given configuration.
func NewRaftCluster(cluster Cluster) *Node {
	node := &Node{
		State:      CandidateState,
		LogChannel: make(chan message.LogData),
		PersistentState: &PersistentState{
			CurrentTerm: 0,
			VotedFor:    nil,
			SelfID:      cluster.OwnID(),
			PeerIPs:     cluster.Nodes(),
		},
		VolatileState: &VolatileState{
			CommitIndex: -1,
			LastApplied: -1,
		},
		VolatileStateLeader: &VolatileStateLeader{},
	}
	return node
}

// randomTicker returns tickers ranging from 150ms to 300ms.
func randomTicker() *time.Ticker {
	randomInt := rand.Intn(150) + 150
	ticker := time.NewTicker(time.Duration(randomInt) * time.Millisecond)
	return ticker
}

// processIncomingData is responsible for parsing the incoming data and calling
// appropriate functions based on the request type.
func processIncomingData(data *incomingData, node *Node) (err error) {

	ctx := context.Background()

	switch data.msg.Kind() {
	case message.KindRequestVoteRequest:
		requestVoteRequest := data.msg.(*message.RequestVoteRequest)
		requestVoteResponse := RequestVoteResponse(node, requestVoteRequest)
		var payload []byte
		payload, err = message.Marshal(requestVoteResponse)
		if err != nil {
			return
		}
		err = data.conn.Send(ctx, payload)
		if err != nil {
			return
		}
	case message.KindAppendEntriesRequest:
		appendEntriesRequest := data.msg.(*message.AppendEntriesRequest)
		appendEntriesResponse := AppendEntriesResponse(node, appendEntriesRequest)
		var payload []byte
		payload, err = message.Marshal(appendEntriesResponse)
		if err != nil {
			return
		}
		err = data.conn.Send(ctx, payload)
		if err != nil {
			return
		}
	}
	return
}
