package message

//go:generate protoc --go_out=. append_entries.proto

var _ Message = (*AppendEntriesRequest)(nil)

// NewLogAppendRequest creates a new append-entries-request message with the
// given parameters.
func NewLogAppendRequest(data string) *LogAppendRequest {
	return &LogAppendRequest{
		Data: data,
	}
}

// Kind returns KindAppendEntriesResponse.
func (*LogAppendRequest) Kind() Kind {
	return KindAppendEntriesResponse
}
