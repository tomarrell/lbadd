package raft

//go:generate stringer -type=State

// State is a raft state that a node can be in.
type State uint8

// Available states
const (
	StateUnknown State = iota
	StateLeader
	StateCandidate
	StateFollower
)
