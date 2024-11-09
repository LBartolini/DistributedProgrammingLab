package protocol

type GetCommand struct {
	writer   *MessageWriter
	filename string
}

func NewGetCommand(writer *MessageWriter, filename string) Command {
	return &GetCommand{writer, filename}
}

func (r *GetCommand) ExecuteAndSend() (n int, err error) {
	return 0, nil
}
