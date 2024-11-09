package protocol

import "os/exec"

type RmCommand struct {
	writer   *MessageWriter
	filename string
}

func NewRmCommand(writer *MessageWriter, filename string) Command {
	return &RmCommand{writer, filename}
}

func (r *RmCommand) ExecuteAndSend() (n int, err error) {
	out, err := exec.Command("rm", BASE_PATH+r.filename).Output()

	if err != nil {
		return 0, err
	}

	return r.writer.Write(out)
}
