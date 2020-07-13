package node

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/engine"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

// Node is a database node. It uses an underlying raft.Server to communicate
// with other nodes, if any.
type Node struct {
	log    zerolog.Logger
	engine engine.Engine

	raft    raft.Server
	cluster cluster.Cluster
}

// New creates a new node that is executing commands on the given executor.
func New(log zerolog.Logger) *Node {
	return &Node{
		log: log,
	}
}

// ListenAndServe starts the node on the given address. The given context must
// be used to stop the server, since there is no stop function. Canceling the
// context or a context timeout will cause the server to attempt a graceful
// shutdown.
func (n *Node) ListenAndServe(ctx context.Context, addr string) error {
	n.log.Info().
		Str("addr", addr).
		Msg("listen and serve")
	return fmt.Errorf("unimplemented")
}

// // Open opens a new cluster, making this node the only node in the cluster.
// // Other clusters can connect to the given address and perform the implemented
// // handshake, in order to become nodes in the cluster.
// func (n *Node) Open(ctx context.Context, addr string) error {
// 	n.log.Info().
// 		Str("addr", addr).
// 		Msg("open")

// 	if err := n.openCluster(ctx, addr); err != nil {
// 		return fmt.Errorf("open cluster: %w", err)
// 	}

// 	return n.startNode()
// }

// // Close closes the node, starting with the underlying raft server, then the
// // cluster, then the executor.
// func (n *Node) Close() error {
// 	ctx := context.TODO()
// 	errs, _ := errgroup.WithContext(ctx)
// 	errs.Go(n.raft.Close)
// 	errs.Go(n.cluster.Close)
// 	errs.Go(n.engine.Close)
// 	return errs.Wait()
// }

func (n *Node) openCluster(ctx context.Context, addr string) error {
	if n.cluster != nil {
		return ErrOpen
	}

	cluster := cluster.NewTCPCluster(n.log)
	cluster.OnConnect(n.performLogonHandshake)
	if err := cluster.Open(ctx, addr); err != nil {
		return fmt.Errorf("open cluster: %w", err)
	}
	return nil
}

func (n *Node) performLogonHandshake(cluster cluster.Cluster, conn network.Conn) {
	n.log.Debug().
		Str("conn-id", conn.RemoteID().String()).
		Msg("perform handshake")

	n.log.Info().
		Str("conn-id", conn.RemoteID().String()).
		Msg("connected")

	cluster.AddConnection(conn)
}

func (n *Node) startNode() error {
	n.raft = raft.NewServer(n.log, n.cluster)
	n.raft.OnReplication(n.replicate)

	return n.raft.Start()
}

// replicate returns the number of commands that were executed from the
// given slice of commands. -1 is returned if no commands were executed.
func (n *Node) replicate(input []*message.Command) int {
	for i := range input {
		cmd := message.ConvertMessageToCommand(input[i])

		// Link to the engine's executor must be added here.
		_, err := n.engine.Evaluate(cmd)
		if err != nil {
			n.log.Error().
				Err(err).
				Msg("failed to replicate input: execute")
			return i - 1
		}
	}
	return len(input)
}
