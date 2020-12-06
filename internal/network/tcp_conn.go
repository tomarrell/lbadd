package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/tomarrell/lbadd/internal/id"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

const (
	frameSizeBytes int = 4
)

var (
	byteOrder = binary.BigEndian
)

var _ Conn = (*tcpConn)(nil)

type tcpConn struct {
	remoteID id.ID
	closed   int32

	readLock   sync.Mutex
	writeLock  sync.Mutex
	underlying net.Conn
}

// DialTCP dials to the given address, assuming a TCP network. The returned Conn
// is ready to use.
func DialTCP(ctx context.Context, ownID id.ID, addr string) (Conn, error) {
	// dial the remote endpoint
	var d net.Dialer
	conn, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("dial tcp: %w", err)
	}

	// create a new connection object
	tcpConn := newTCPConn(conn)

	// receive the remote ID from the remote endpoint and apply it
	remoteID, err := tcpConn.Receive(ctx)
	if err != nil {
		_ = tcpConn.Close()
		return nil, fmt.Errorf("receive remote ID: %w", err)
	}
	parsedID, err := id.Parse(remoteID)
	if err != nil {
		_ = tcpConn.Close()
		return nil, fmt.Errorf("parse remote ID: %w", err)
	}
	tcpConn.remoteID = parsedID

	// send own ID to remote endpoint
	err = tcpConn.Send(ctx, ownID.Bytes())
	if err != nil {
		_ = tcpConn.Close()
		return nil, fmt.Errorf("send own ID: %w", err)
	}

	// return the connection object
	return tcpConn, nil
}

// NewTCPConn wraps the underlying connection into a tcpConn.
func NewTCPConn(underlying net.Conn) Conn {
	return newTCPConn(underlying)
}

func newTCPConn(underlying net.Conn) *tcpConn {
	id := id.Create()
	conn := &tcpConn{
		remoteID:   id,
		underlying: underlying,
	}
	return conn
}

func (c *tcpConn) RemoteID() id.ID {
	return c.remoteID
}

func (c *tcpConn) Send(ctx context.Context, payload []byte) error {
	if atomic.LoadInt32(&c.closed) == 1 {
		return ErrClosed
	}

	c.writeLock.Lock()
	defer c.writeLock.Unlock()

	if deadline, ok := ctx.Deadline(); ok {
		// Set the write deadline on the underlying connection according to the
		// given context. This write deadline applies to the whole function, so
		// we only set it once here. On the next write-call, it will be set
		// again, or will be reset in the else block, to not keep an old
		// deadline.
		_ = c.underlying.SetWriteDeadline(deadline)
	} else {
		_ = c.underlying.SetWriteDeadline(time.Time{}) // remove the write deadline
	}

	select {
	case err := <-c.sendAsync(payload):
		return err
	case <-ctx.Done():
		return ErrTimeout
	}
}

func (c *tcpConn) sendAsync(payload []byte) chan error {
	result := make(chan error)
	go func() {
		var frameSize [frameSizeBytes]byte
		byteOrder.PutUint32(frameSize[:], uint32(len(payload)))

		n, err := c.underlying.Write(frameSize[:])
		if err != nil {
			// if the error is a timeout, yield a plain timeout error
			if netErr, ok := err.(*net.OpError); ok && netErr.Timeout() {
				result <- ErrTimeout
				return
			}
			result <- fmt.Errorf("write size: %w", err)
			return
		}
		if n != frameSizeBytes {
			result <- fmt.Errorf("write bytes: written %v of %v size bytes", n, len(payload))
			return
		}

		n, err = c.underlying.Write(payload)
		if err != nil {
			result <- fmt.Errorf("write payload: %w", err)
			return
		}
		if n != len(payload) {
			result <- fmt.Errorf("write bytes: written %v of %v payload bytes", n, len(payload))
			return
		}

		result <- nil
	}()
	return result
}

func (c *tcpConn) Receive(ctx context.Context) ([]byte, error) {
	if atomic.LoadInt32(&c.closed) == 1 {
		return nil, ErrClosed
	}

	c.readLock.Lock()
	defer c.readLock.Unlock()

	if deadline, ok := ctx.Deadline(); ok {
		// Set the read deadline on the underlying connection according to the
		// given context. This read deadline applies to the whole function, so
		// we only set it once here. On the next read-call, it will be set
		// again, or will be reset in the else block, to not keep an old
		// deadline.
		_ = c.underlying.SetReadDeadline(deadline)
	} else {
		_ = c.underlying.SetReadDeadline(time.Time{}) // remove the read deadline
	}

	select {
	case res := <-c.receiveAsync():
		if err, ok := res.(error); ok {
			return nil, err
		}
		return res.([]byte), nil
	case <-ctx.Done():
		return nil, ErrTimeout
	}
}

func (c *tcpConn) receiveAsync() chan interface{} {
	result := make(chan interface{})
	go func() {
		var frameSizeB [frameSizeBytes]byte
		n, err := c.underlying.Read(frameSizeB[:])
		if err != nil {
			// if the error is a timeout, yield a plain timeout error
			if netErr, ok := err.(*net.OpError); ok && netErr.Timeout() {
				result <- ErrTimeout
				return
			}
			result <- fmt.Errorf("read frame size: %w", err)
			return
		}
		if n != frameSizeBytes {
			result <- fmt.Errorf("read only %v frame size bytes of %v expected", n, frameSizeBytes)
			return
		}

		frameSize := byteOrder.Uint32(frameSizeB[:])
		frameData := make([]byte, frameSize)
		n, err = c.underlying.Read(frameData)
		if err != nil {
			result <- fmt.Errorf("read frame payload: %w", err)
			return
		}
		if n != int(frameSize) {
			result <- fmt.Errorf("read only %v frame payload bytes of %v expected", n, frameSize)
			return
		}
		result <- frameData
	}()
	return result
}

func (c *tcpConn) Close() error {
	atomic.StoreInt32(&c.closed, 1)

	// release all resources
	ctx := context.Background()
	errs, _ := errgroup.WithContext(ctx)
	errs.Go(c.underlying.Close)
	return errs.Wait()
}
