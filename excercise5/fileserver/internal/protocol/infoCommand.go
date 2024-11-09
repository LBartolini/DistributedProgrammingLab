package protocol

import "os/exec"

type InfoCommand struct {
	writer   *MessageWriter
	filename string
}

func NewInfoCommand(writer *MessageWriter, filename string) Command {
	return &InfoCommand{writer, filename}
}

func (r *InfoCommand) ExecuteAndSend() (n int, err error) {
	out, err := exec.Command("stat", BASE_PATH+r.filename).Output()

	if err != nil {
		return 0, err
	}

	return r.writer.Write(out)
}
