package network

import (
	"github.com/rs/zerolog"
)

// TestNode describes a single node and all its connections.
type TestNode struct {
	Node  Server
	Conns []Conn
}

// NewTestNode returns a ready to use node with no connections associated.
func NewTestNode(IP, port string, log zerolog.Logger) (tNode *TestNode, err error) {
	node := NewTCPServer(log)
	go func() {
		err = node.Open(IP + ":" + port)
	}()

	if err != nil {
		return
	}

	<-node.Listening()
	return &TestNode{
		Node: node,
	}, nil
}
