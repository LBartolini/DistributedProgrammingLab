package protocol

import (
	"io"
)

type MessageWriter struct {
	writer io.Writer
}

func NewMessageWriter(writer io.Writer) *MessageWriter {
	return &MessageWriter{writer}
}
