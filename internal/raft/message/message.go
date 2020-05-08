package message

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"google.golang.org/protobuf/proto"
)

type Message interface {
	proto.Message
	Kind() Kind
}

func Marshal(m Message) ([]byte, error) {
	data, err := proto.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("proto marshal: %w", err)
	}

	var buf bytes.Buffer
	buf.WriteByte(byte(m.Kind()))
	buf.Write(data)
	return buf.Bytes(), nil
}

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
