package raft

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
	networkmocks "github.com/tomarrell/lbadd/internal/network/mocks"
	"github.com/tomarrell/lbadd/internal/raft/mocks"
)

func TestExample(t *testing.T) {
	conn := new(networkmocks.Conn)
	// mock call to RemoteID
	id := id.Create()
	conn.
		On("RemoteID").
		Return(id)

	c := new(mocks.Cluster)
	// mock call to Nodes
	c.
		On("Nodes").
		Return([]network.Conn{conn})

	assert.Equal(t, conn.RemoteID().String(), c.Nodes()[0].RemoteID().String())
}
