package message

//go:generate stringer -type=Kind

// Kind describes a kind of a message, that is used by the raft module.
type Kind uint32

// Available kinds
const (
	// KindUnknown must not be used. It is the default value for Kind. If this
	// value occurs, something was not initialized properly.
	KindUnknown Kind = iota

	KindAppendEntriesRequest
	KindAppendEntriesResponse

	KindFollowerLocationListRequest
	KindFollowerLocationListResponse

	KindLeaderLocationRequest
	KindLeaderLocationResponse

	KindRequestVoteRequest
	KindRequestVoteResponse
)
