package network

import (
	"encoding/binary"
	"fmt"
	"net"

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
	id         ID
	underlying net.Conn
	closed     bool
}

// DialTCP dials to the given address, assuming a TCP network. The returned Conn
// is ready to use.
func DialTCP(addr string) (Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("dial tcp: %w", err)
	}
	tcpConn := newTCPConn(conn)
	myID, err := tcpConn.Receive()
	if err != nil {
		_ = tcpConn.Close()
		return nil, fmt.Errorf("receive ID: %w", err)
	}
	parsedID, err := parseID(myID)
	if err != nil {
		_ = tcpConn.Close()
		return nil, fmt.Errorf("parse ID: %w", err)
	}
	tcpConn.id = parsedID
	return tcpConn, nil
}

func newTCPConn(underlying net.Conn) *tcpConn {
	id := createID()
	conn := &tcpConn{
		id:         id,
		underlying: underlying,
	}
	return conn
}

func (c *tcpConn) ID() ID {
	return c.id
}

func (c *tcpConn) Send(payload []byte) error {
	if c.closed {
		return ErrClosed
	}

	var frameSize [frameSizeBytes]byte
	byteOrder.PutUint32(frameSize[:], uint32(len(payload)))

	n, err := c.underlying.Write(frameSize[:])
	if err != nil {
		return fmt.Errorf("write size: %w", err)
	}
	if n != frameSizeBytes {
		return fmt.Errorf("write bytes: written %v of %v size bytes", n, len(payload))
	}

	n, err = c.underlying.Write(payload)
	if err != nil {
		return fmt.Errorf("write payload: %w", err)
	}
	if n != len(payload) {
		return fmt.Errorf("write bytes: written %v of %v payload bytes", n, len(payload))
	}
	return nil
}

func (c *tcpConn) Receive() ([]byte, error) {
	var frameSizeB [frameSizeBytes]byte
	n, err := c.underlying.Read(frameSizeB[:])
	if err != nil {
		return nil, fmt.Errorf("read frame size: %w", err)
	}
	if n != frameSizeBytes {
		return nil, fmt.Errorf("read only %v frame size bytes of %v expected", n, frameSizeBytes)
	}

	frameSize := byteOrder.Uint32(frameSizeB[:])
	frameData := make([]byte, frameSize)
	n, err = c.underlying.Read(frameData)
	if err != nil {
		return nil, fmt.Errorf("read frame payload: %w", err)
	}
	if n != int(frameSize) {
		return nil, fmt.Errorf("read only %v frame payload bytes of %v expected", n, frameSize)
	}
	return frameData, nil
}

func (c *tcpConn) Close() error {
	c.closed = true

	// release all resources
	ctx := context.Background()
	errs, _ := errgroup.WithContext(ctx)
	errs.Go(c.underlying.Close)
	return errs.Wait()
}
