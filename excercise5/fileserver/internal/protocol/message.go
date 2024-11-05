package protocol

import (
	"errors"
)

type Message interface {
	Send(writer *MessageWriter) error
}

type Command struct {
	operation string
	fileName  string
}

//-------

type lsCommand struct {
	command Command
}

func (r *lsCommand) Send(writer *MessageWriter) error {
	return nil
}

//-------

type catCommand struct {
	command Command
}

func (r *catCommand) Send(writer *MessageWriter) error {
	return nil
}

//-------

type rmCommand struct {
	command Command
}

func (r *rmCommand) Send(writer *MessageWriter) error {
	return nil
}

//-------

type getCommand struct {
	command Command
}

func (r *getCommand) Send(writer *MessageWriter) error {
	return nil
}

//-------

type infoCommand struct {
	command Command
}

func (r *infoCommand) Send(writer *MessageWriter) error {
	return nil
}

var UnknownCommand = errors.New("Unknown command")
