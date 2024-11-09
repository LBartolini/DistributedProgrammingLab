package protocol

import (
	"fmt"
	"io"
)

type MessageWriter struct {
	writer io.Writer
}

func NewMessageWriter(writer io.Writer) *MessageWriter {
	return &MessageWriter{writer}
}

func (w *MessageWriter) Write(p []byte) (n int, err error) {
	return fmt.Fprintf(w.writer, "%s\n>>>", p)
}
