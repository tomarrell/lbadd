package node

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/compile"
	"github.com/tomarrell/lbadd/internal/executor"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
	"github.com/tomarrell/lbadd/internal/raft/message"
	"golang.org/x/sync/errgroup"
)

// Node is a database node. It uses an underlying raft.Server to communicate
// with other nodes, if any.
type Node struct {
	log  zerolog.Logger
	exec executor.Executor

	raft    raft.Server
	cluster cluster.Cluster
}

// New creates a new node that is executing commands on the given executor.
func New(log zerolog.Logger, exec executor.Executor) *Node {
	return &Node{
		log:  log,
		exec: exec,
	}
}

// Open opens a new cluster, making this node the only node in the cluster.
// Other clusters can connect to the given address and perform the implemented
// handshake, in order to become nodes in the cluster.
func (n *Node) Open(ctx context.Context, addr string) error {
	n.log.Info().
		Str("addr", addr).
		Msg("open")

	if err := n.openCluster(ctx, addr); err != nil {
		return fmt.Errorf("open cluster: %w", err)
	}

	return n.startNode()
}

// Close closes the node, starting with the underlying raft server, then the
// cluster, then the executor.
func (n *Node) Close() error {
	ctx := context.TODO()
	errs, _ := errgroup.WithContext(ctx)
	errs.Go(n.raft.Close)
	errs.Go(n.cluster.Close)
	errs.Go(n.exec.Close)
	return errs.Wait()
}

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

func (n *Node) replicate(input []*message.Command) int {
	for i := range input {
		cmd := convert(input[i])

		res, err := n.exec.Execute(cmd)
		if err != nil {
			n.log.Error().
				Err(err).
				Msg("failed to replicate input: execute")
			return 0
		}

		_ = res // ignore the result, because we don't need it to be printed or processed anywhere
	}
	// TODO - return appropriate values of executed commands.
	return -1
}

// convert is a stop gap arrangement until the compile.Command aligns with the universal format for IR commands.
func convert(input *message.Command) compile.Command {
	return compile.Command{}
}
