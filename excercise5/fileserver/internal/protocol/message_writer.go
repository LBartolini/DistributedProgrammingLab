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

func (w *MessageWriter) Write(command Command) (n int, err error) {
	return command.ExecuteAndSend(w)
}
