package raft

import "github.com/rs/zerolog"

// SimpleRaftTest implements TestFramework.
type SimpleRaftTest struct {
	log        zerolog.Logger
	parameters OperationParameters
	config     NetworkConfiguration
}

// BeginTest will wrapped under a Go Test for ease of use.
func (t *SimpleRaftTest) BeginTest() {
	// Check for proper config before beginning.
}
