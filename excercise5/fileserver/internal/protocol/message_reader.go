package protocol

import (
	"bufio"
	"io"
	"strings"
)

// MessageReader allows to deserialize a Message from a stream
type MessageReader struct {
	reader bufio.Reader
}

// NewMessageReader returns a new MessageReader that gets data from an underlying reader.
func NewMessageReader(reader io.Reader) *MessageReader {
	return &MessageReader{*bufio.NewReader(reader)}
}

// Read deserializes a Message from a stream.
func (r *MessageReader) Read() (Message, error) {
	messageType, err := r.reader.ReadString(' ')

	messageType = strings.TrimSpace(messageType)

	if err != nil {
		return nil, err
	}

	switch messageType {
	case "SEND":
		content, err := r.reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		content = strings.TrimSpace(content)

		return SendMessage{content}, nil
	case "NAME":
		newName, err := r.reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		newName = strings.TrimSpace(newName)

		return ChangeNameMessage{newName}, nil
	case "MESSAGE":
		author, err := r.reader.ReadString(' ')

		if err != nil {
			return nil, err
		}

		author = strings.TrimSpace(author)

		content, err := r.reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		content = strings.TrimSpace(content)

		return NotifyMessage{author, content}, nil

	default:
		return nil, UnknownCommand
	}
}
