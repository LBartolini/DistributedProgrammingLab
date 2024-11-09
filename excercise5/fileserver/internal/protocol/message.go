package protocol

import (
	"errors"
)

const BASE_PATH = "dbox/"

type Command interface {
	ExecuteAndSend() (n int, err error)
}

var ErrUnknownCommand = errors.New("unknown command")
var ErrNotEnoughArguments = errors.New("uot enough arguments for command provided")
