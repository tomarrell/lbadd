package message

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"google.golang.org/protobuf/proto"
)

// Message describes a serializable, more or less self-describing protobuf
// message. A message consists of a kind (message.Kind) and an actual protobuf
// message.
type Message interface {
	proto.Message
	// Kind returns the kind of this message. If this returns
	// message.KindUnknown, something went wrong, or the client and server
	// versions are not matching.
	Kind() Kind
}

// Marshal converts the given message to a byte slice that can be unmarshalled
// with message.Unmarshal. The kind is encoded with 4 bytes big endian as uint32
// and is the first 4 bytes of the serialized message. The rest of the message
// is the serialized protobuf message.
func Marshal(m Message) ([]byte, error) {
	data, err := proto.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("proto marshal: %w", err)
	}

	var buf bytes.Buffer
	kind := make([]byte, 4)
	binary.BigEndian.PutUint32(kind, uint32(m.Kind()))
	buf.Write(kind)
	buf.Write(data)
	return buf.Bytes(), nil
}

// Unmarshal converts bytes to a message. For using the returned message, check
// Message.Kind() and process a it accordingly.
func Unmarshal(data []byte) (Message, error) {
	kindBytes := data[:4] // kind is uint32, which has 4 bytes
	payload := data[4:]

	kind := Kind(binary.BigEndian.Uint32(kindBytes))
	var msg Message
	switch kind {
	case KindRequestVoteRequest:
		msg = &RequestVoteRequest{}
	case KindRequestVoteResponse:
		msg = &RequestVoteResponse{}
	default:
		return nil, ErrUnknownKind
	}

	if err := proto.Unmarshal(payload, msg); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return msg, nil
}
