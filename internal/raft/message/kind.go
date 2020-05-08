package message

//go:generate stringer -type=Kind

type Kind uint32

const (
	KindUnknown Kind = iota

	KindAppendEntriesRequest
	KindAppendEntriesResponse

	KindRequestVoteRequest
	KindRequestVoteResponse
)
