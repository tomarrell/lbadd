package node

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/compile"
	"github.com/tomarrell/lbadd/internal/executor"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/parser"
	"github.com/tomarrell/lbadd/internal/raft"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
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

func (n *Node) Open(ctx context.Context, addr string) error {
	n.log.Info().
		Str("addr", addr).
		Msg("open")

	if err := n.openCluster(ctx, addr); err != nil {
		return fmt.Errorf("open cluster: %w", err)
	}

	return n.startNode()
}

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
		Str("conn-id", conn.ID().String()).
		Msg("perform handshake")

	n.log.Info().
		Str("conn-id", conn.ID().String()).
		Msg("connected")

	cluster.AddConnection(conn)
}

func (n *Node) startNode() error {
	n.raft = raft.NewServer(n.cluster)
	n.raft.OnReplication(n.replicate)

	return n.raft.Start()
}

func (n *Node) replicate(input string) {
	parser := parser.New(input)
	for {
		stmt, errs, ok := parser.Next()
		if !ok {
			break // no more statements
		}
		if len(errs) != 0 {
			// if errors occur, abort replication of this input, even if there
			// may be correct statements in the input
			logErrs := zerolog.Arr()
			for _, err := range errs {
				logErrs.Err(err)
			}
			n.log.Error().
				Array("errors", logErrs).
				Msg("failed to replicate input: parse")
			return
		}

		compiler := compile.NewSimpleCompiler()
		cmd, err := compiler.Compile(stmt)
		if err != nil {
			n.log.Error().
				Err(err).
				Msg("failed to replicate input: compile")
			return
		}

		res, err := n.exec.Execute(cmd)
		if err != nil {
			n.log.Error().
				Err(err).
				Msg("failed to replicate input: execute")
			return
		}

		_ = res // ignore the result, because we don't need it to be printed or processed anywhere
	}
}
