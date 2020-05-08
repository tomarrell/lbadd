package network

import (
	"fmt"
	"net"
	"time"

	"github.com/rs/zerolog"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

var _ Server = (*tcpServer)(nil)

type tcpServer struct {
	log zerolog.Logger

	open      bool
	listening chan struct{}
	lis       net.Listener

	onConnect ConnHandler
}

// NewTCPServer creates a new ready to use TCP server that uses the given
// logger.
func NewTCPServer(log zerolog.Logger) Server {
	return &tcpServer{
		log:       log,
		listening: make(chan struct{}),
	}
}

func (s *tcpServer) Open(addr string) error {
	if s.open {
		return ErrOpen
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	s.open = true
	s.lis = lis

	s.log.Debug().
		Str("addr", lis.Addr().String()).
		Msg("tcp open")

	s.handleIncomingConnections()
	return nil
}

func (s *tcpServer) Listening() <-chan struct{} {
	return s.listening
}

func (s *tcpServer) Addr() net.Addr {
	if s.lis == nil {
		return nil
	}
	return s.lis.Addr()
}

func (s *tcpServer) OnConnect(h ConnHandler) {
	s.onConnect = h
}

func (s *tcpServer) Close() error {
	s.open = false

	// release all resources
	ctx := context.Background()
	errs, _ := errgroup.WithContext(ctx)
	errs.Go(s.lis.Close)
	return errs.Wait()
}

func (s *tcpServer) handleIncomingConnections() {
	close(s.listening)
	for {
		conn, err := s.lis.Accept()
		if err != nil {
			if !s.open {
				// server was already closed, we can discard the error, but we
				// also need to stop accepting further connections
				break
			}

			// otherwise, an error occurred while accepting a connection, but
			// since we can't that let us stop from accepting further
			// connections, we just log it
			s.log.Error().
				Err(err).
				Msg("accept")
			continue
		}

		go s.handleIncomingNetConn(conn)
	}
}

func (s *tcpServer) handleIncomingNetConn(conn net.Conn) {
	tcpConn := newTCPConn(conn)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := tcpConn.Send(ctx, tcpConn.id.Bytes())
	if err != nil {
		s.log.Error().
			Err(err).
			Msg("send ID")
		_ = tcpConn.Close()
		return
	}

	if s.onConnect != nil {
		s.onConnect(tcpConn)
	}
}
