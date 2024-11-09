package protocol

import (
	"os/exec"
)

type LsCommand struct {
	writer *MessageWriter
}

func NewLsCommand(writer *MessageWriter) Command {
	return &LsCommand{writer}
}

func (r *LsCommand) ExecuteAndSend() (n int, err error) {
	out, err := exec.Command("ls", BASE_PATH).Output()

	if err != nil {
		return 0, err
	}

	return r.writer.Write(out)
}
