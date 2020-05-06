package network

import (
	"fmt"
	"net"

	"github.com/rs/zerolog"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

var _ Server = (*tcpServer)(nil)

type tcpServer struct {
	log zerolog.Logger

	open bool
	lis  net.Listener

	onConnect ConnHandler
}

// NewTCPServer creates a new ready to use TCP server that uses the given
// logger.
func NewTCPServer(log zerolog.Logger) Server {
	return &tcpServer{
		log: log,
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
		}

		tcpConn := newTCPConn(conn)
		tcpConn.Send(tcpConn.id.Bytes())

		if s.onConnect != nil {
			go s.onConnect(tcpConn)
		}
	}
}
