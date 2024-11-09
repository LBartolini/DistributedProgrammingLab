package protocol

import "os/exec"

type CatCommand struct {
	writer   *MessageWriter
	filename string
}

func NewCatCommand(writer *MessageWriter, filename string) Command {
	return &CatCommand{writer, filename}
}

func (r *CatCommand) ExecuteAndSend() (n int, err error) {
	out, err := exec.Command("cat", BASE_PATH+r.filename).Output()

	if err != nil {
		return 0, err
	}

	return r.writer.Write(out)
}
