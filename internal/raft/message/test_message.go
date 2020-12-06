package message

//go:generate protoc --go_out=. test_message.proto

var _ Message = (*TestMessage)(nil)

// NewTestMessage creates a new test message with the given data.
func NewTestMessage(data string) *TestMessage {
	return &TestMessage{
		Data: data,
	}
}

// Kind returns KindTestMessage.
func (*TestMessage) Kind() Kind {
	return KindTestMessage
}
