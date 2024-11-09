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
func (r *MessageReader) Read(writer *MessageWriter) (Command, error) {
	message, err := r.reader.ReadString('\n')

	fields := strings.Fields(message)

	if err != nil || len(fields) == 0 {
		return nil, err
	}

	messageType := fields[0]
	fields = fields[1:]

	switch messageType {
	case "ls":
		return NewLsCommand(writer), nil
	case "cat":
		if len(fields) == 0 {
			return nil, ErrNotEnoughArguments
		}

		filename := strings.TrimSpace(fields[0])

		return NewCatCommand(writer, filename), nil
	case "rm":
		if len(fields) == 0 {
			return nil, ErrNotEnoughArguments
		}

		filename := strings.TrimSpace(fields[0])

		return NewRmCommand(writer, filename), nil

	case "get":
		if len(fields) == 0 {
			return nil, ErrNotEnoughArguments
		}

		filename := strings.TrimSpace(fields[0])

		return NewGetCommand(writer, filename), nil

	case "info":
		if len(fields) == 0 {
			return nil, ErrNotEnoughArguments
		}

		filename := strings.TrimSpace(fields[0])

		return NewInfoCommand(writer, filename), nil

	default:
		return nil, ErrUnknownCommand
	}
}
