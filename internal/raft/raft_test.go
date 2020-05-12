package raft

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
)

func Test_NewServer(t *testing.T) {
	assert := assert.New(t)

	log := zerolog.Nop()
	cluster := cluster.NewTCPCluster(log)
	server := NewServer(
		log,
		cluster,
	)
	err := server.Start()
	assert.NoError(err)
}
