package cluster

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/message"
	"golang.org/x/sync/errgroup"
)

const (
	tcpClusterMessageQueueBufferSize = 5
)

var _ Cluster = (*tcpCluster)(nil)

type tcpCluster struct {
	log zerolog.Logger

	connLock sync.Mutex
	conns    []network.Conn

	onConnect ConnHandler

	server   network.Server
	messages chan incomingPayload
	started  chan struct{}
	closed   int32
}

type incomingPayload struct {
	origin  network.Conn
	payload []byte
}

// NewTCPCluster creates a new cluster that uses TCP connections to communicate
// with other nodes.
func NewTCPCluster(log zerolog.Logger) Cluster {
	return &tcpCluster{
		log:       log.With().Str("cluster", "tcp").Logger(),
		onConnect: func(c Cluster, conn network.Conn) { c.AddConnection(conn) },
		server:    network.NewTCPServer(log),
		messages:  make(chan incomingPayload, tcpClusterMessageQueueBufferSize),
		started:   make(chan struct{}),
	}
}

func (c *tcpCluster) OnConnect(handler ConnHandler) {
	c.onConnect = handler
}

func (c *tcpCluster) Join(ctx context.Context, addr string) error {
	// connect to the given address
	conn, err := network.DialTCP(ctx, c.server.OwnID(), addr)
	if err != nil {
		return fmt.Errorf("dial tcp: %w", err)
	}
	c.AddConnection(conn)

	// We have now joined the cluster, start the common procedure for network
	// operations, like listening to incoming connections, messages etc.
	go c.start()

	return nil
}

func (c *tcpCluster) Open(ctx context.Context, addr string) error {
	go func() {
		_ = c.server.Open(addr)
	}()

	select {
	case <-ctx.Done():
		_ = c.Close() // will also close the server that we just tried to open
		return ErrTimeout
	case <-c.server.Listening():
	}
	go c.start()
	return nil
}

// Nodes returns a copy of the connections that the cluster currently holds.
func (c *tcpCluster) Nodes() []network.Conn {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	nodes := make([]network.Conn, len(c.conns))
	copy(nodes, c.conns)
	return nodes
}

func (c *tcpCluster) OwnID() id.ID {
	return c.server.OwnID()
}

func (c *tcpCluster) Receive(ctx context.Context) (network.Conn, message.Message, error) {
	incoming, ok := <-c.messages
	if !ok {
		return nil, nil, fmt.Errorf("channel closed")
	}
	msg, err := message.Unmarshal(incoming.payload)
	if err != nil {
		return nil, nil, fmt.Errorf("unmarshal: %w", err)
	}
	return incoming.origin, msg, nil
}

func (c *tcpCluster) Broadcast(ctx context.Context, msg message.Message) error {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	errs, _ := errgroup.WithContext(ctx)
	for _, conn := range c.conns {
		errs.Go(func() error {
			if err := c.sendMessage(ctx, conn, msg); err != nil {
				return fmt.Errorf("send message: %w", err)
			}
			return nil
		})
	}
	return errs.Wait()
}

// Close will shut down the cluster. This means:
//
//  * the cluster's status is set to closed
//  * all connections in the cluster's connection list are closed (not removed)
//  * the underlying network server is closed
//  * the cluster's message queue is closed
//
// After Close is called on this cluster, it is no longer usable.
func (c *tcpCluster) Close() error {
	atomic.StoreInt32(&c.closed, 1)

	// close all connections
	var errs errgroup.Group
	c.connLock.Lock()
	for _, conn := range c.conns {
		errs.Go(conn.Close)
	}
	c.connLock.Unlock()

	errs.Go(c.server.Close)

	// close the message queue
	close(c.messages)

	return errs.Wait()
}

// addConnection will add the connection to the list of connections of this
// cluster. It will also start a goroutine that reads from the connection. That
// goroutine will push back read data.
func (c *tcpCluster) AddConnection(conn network.Conn) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	c.conns = append(c.conns, conn)
	go c.receiveMessages(conn)
}

// RemoveConnection will attempt to remove the given connection from the list of
// connections in this cluster. If the connection was found, it will be removed
// AND CLOSED. If the connection was NOT found, it will NOT be closed.
func (c *tcpCluster) RemoveConnection(conn network.Conn) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	for i, node := range c.conns {
		if node.RemoteID() == conn.RemoteID() {
			c.conns[i] = c.conns[len(c.conns)-1]
			c.conns[len(c.conns)-1] = nil
			c.conns = c.conns[:len(c.conns)-1]

			_ = conn.Close()
			return
		}
	}
}

func (c *tcpCluster) sendMessage(ctx context.Context, conn network.Conn, msg message.Message) error {
	msgData, err := message.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	if err := conn.Send(ctx, msgData); err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

func (c *tcpCluster) start() {
	// On connect, execute the on-connect hook.
	c.server.OnConnect(func(conn network.Conn) {
		if c.onConnect != nil {
			c.onConnect(c, conn)
		}
	})

	// signal all waiting receive message goroutines that the server is now
	// started and they can start pushing messages onto the queue
	close(c.started)
}

// receiveMessages will wait for the cluster to be started, and then, while the
// cluster is not closed, attempt to read data from the connection. If the read
// times out, it tries again indefinitely. If an error occurs during the read,
// and the server is already closed, nothing happens, but this method returns.
// If an error occurs during the read, and the server is NOT closed, the
// connection will be removed with (*tcpCluster).RemoveConnection, and the error
// will be logged with error level. After that, this method will return.
func (c *tcpCluster) receiveMessages(conn network.Conn) {
	<-c.started // wait for the server to be started

	for atomic.LoadInt32(&c.closed) == 0 {
		// receive data from the connection
		data, err := conn.Receive(context.TODO())
		if err != nil {
			if err == network.ErrTimeout {
				// didn't receive a message within the timeout, try again
				continue
			}
			if atomic.LoadInt32(&c.closed) == 1 {
				// server is closed, no reason to log errors from connections
				// that we failed to read from, but break the read loop and
				// terminate this goroutine
				return
			}
			c.RemoveConnection(conn) // also closes the connection
			c.log.Error().
				Err(err).
				Str("fromID", conn.RemoteID().String()).
				Msg("receive failed, removing connection")
			return // abort this goroutine
		}

		// push payload and connection onto the message queue
		c.messages <- incomingPayload{
			origin:  conn,
			payload: data,
		}
	}
}
