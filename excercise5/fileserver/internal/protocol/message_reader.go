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
func (r *MessageReader) Read() (Command, error) {
	messageType, err := r.reader.ReadString(' ')

	messageType = strings.TrimSpace(messageType)

	if err != nil {
		return nil, err
	}

	switch messageType {
	case "ls":
		return NewLsCommand(), nil
	case "cat":
		content, err := r.reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		filename := strings.TrimSpace(content)

		return NewCatCommand(filename), nil
	case "rm":
		content, err := r.reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		filename := strings.TrimSpace(content)

		return NewRmCommand(filename), nil

	case "get":
		content, err := r.reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		filename := strings.TrimSpace(content)

		return NewGetCommand(filename), nil

	case "info":
		content, err := r.reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		filename := strings.TrimSpace(content)

		return NewInfoCommand(filename), nil

	default:
		return nil, UnknownCommand
	}
}
