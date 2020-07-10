package message

import (
	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/id"
)

//go:generate protoc --go_out=. append_entries.proto

var _ Message = (*AppendEntriesRequest)(nil)

// NewAppendEntriesRequest creates a new append-entries-request message with the
// given parameters.
func NewAppendEntriesRequest(term int32, leaderID id.ID, prevLogIndex int32, prevLogTerm int32, entries []*LogData, leaderCommit int32) *AppendEntriesRequest {
	return &AppendEntriesRequest{
		Term:         term,
		LeaderID:     leaderID.Bytes(),
		PrevLogIndex: prevLogIndex,
		PrevLogTerm:  prevLogTerm,
		Entries:      entries,
		LeaderCommit: leaderCommit,
	}
}

// Kind returns KindAppendEntriesRequest.
func (*AppendEntriesRequest) Kind() Kind {
	return KindAppendEntriesRequest
}

// NewLogData creates a new log-data object, which can be used for an
// append-entries-request message.
func NewLogData(term int32, data command.Command) *LogData {
	msg, _ := ConvertCommandToMessage(data)

	return &LogData{
		Term:  term,
		Entry: msg.(*Command),
	}
}

// NewAppendEntriesResponse creates a new append-entries-response message with
// the given parameters.
func NewAppendEntriesResponse(term int32, success bool) *AppendEntriesResponse {
	return &AppendEntriesResponse{
		Term:    term,
		Success: success,
	}
}

// Kind returns KindAppendEntriesResponse.
func (*AppendEntriesResponse) Kind() Kind {
	return KindAppendEntriesResponse
}
