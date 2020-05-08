package message

import (
	"github.com/tomarrell/lbadd/internal/id"
)

//go:generate protoc --go_out=. append_entries.proto

var _ Message = (*AppendEntriesRequest)(nil)

func NewAppendEntriesRequest(term int32, leaderID id.ID, prevLogIndex int32, prevLogTerm int32, entries []*LogData, leaderCommit int32) Message {
	return &AppendEntriesRequest{
		Term:         term,
		LeaderId:     leaderID.Bytes(),
		PrevLogIndex: prevLogIndex,
		PrevLogTerm:  prevLogTerm,
		Entries:      entries,
		LeaderCommit: leaderCommit,
	}
}

func (*AppendEntriesRequest) Kind() Kind {
	return KindAppendEntriesRequest
}

func NewLogData(term int32, data string) *LogData {
	return &LogData{
		Term: term,
		Data: data,
	}
}

func NewAppendEntriesResponse(term int32, success bool) Message {
	return &AppendEntriesResponse{
		Term:    term,
		Success: success,
	}
}

func (*AppendEntriesResponse) Kind() Kind {
	return KindAppendEntriesResponse
}
