package raft

import (
	"context"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
)

// Raft integration tests go here.
func Test_NewServer(t *testing.T) {
	assert := assert.New(t)

	log := zerolog.Nop()
	ctx := context.Background()
	cluster := cluster.NewTCPCluster(log)
	err := cluster.Open(ctx, ":0")
	server := NewServer(
		log,
		cluster,
	)
	assert.NoError(err)
	err = server.Start()
	assert.NoError(err)
}
