package protocol

import (
	"errors"
	"os/exec"
)

const BASE_PATH = "dbox"

type Command interface {
	ExecuteAndSend(writer *MessageWriter) (n int, err error)
}

//-------

type LsCommand struct{}

func NewLsCommand() Command {
	return &LsCommand{}
}

func (r *LsCommand) ExecuteAndSend(writer *MessageWriter) (n int, err error) {
	out, err := exec.Command("ls", BASE_PATH).Output()

	if err != nil {
		return 0, err
	}

	return writer.writer.Write(out)
}

//-------

type CatCommand struct {
	filename string
}

func NewCatCommand(filename string) Command {
	return &CatCommand{filename}
}

func (r *CatCommand) ExecuteAndSend(writer *MessageWriter) (n int, err error) {
	return 0, nil
}

//-------

type RmCommand struct {
	filename string
}

func NewRmCommand(filename string) Command {
	return &RmCommand{filename}
}

func (r *RmCommand) ExecuteAndSend(writer *MessageWriter) (n int, err error) {
	return 0, nil
}

//-------

type GetCommand struct {
	filename string
}

func NewGetCommand(filename string) Command {
	return &GetCommand{filename}
}

func (r *GetCommand) ExecuteAndSend(writer *MessageWriter) (n int, err error) {
	return 0, nil
}

//-------

type InfoCommand struct {
	filename string
}

func NewInfoCommand(filename string) Command {
	return &InfoCommand{filename}
}

func (r *InfoCommand) ExecuteAndSend(writer *MessageWriter) (n int, err error) {
	return 0, nil
}

//-------

var UnknownCommand = errors.New("Unknown command")
