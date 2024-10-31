package gotodo

import (
	"errors"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos struct {
	todos   []Todo
	render  Render
	storage Storage
}

var ErrIndexOutOfBounds = errors.New("invalid index")
var ErrEmptyTitle = errors.New("No description provided")

func NewTodos(render Render, storage Storage) *Todos {
	return &Todos{make([]Todo, 0), render, storage}
}

func (todos *Todos) Add(title string) error {
	if title == "" {
		return ErrEmptyTitle
	}

	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	ts := *todos
	(*todos).todos = append(ts.todos, todo)

	return nil
}

func (todos *Todos) validateIndex(index int) error {
	ts := *todos
	if index < 0 || index >= len(ts.todos) {
		return ErrIndexOutOfBounds
	}
	return nil
}

func (todos *Todos) Delete(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	ts := *todos
	ts.todos = append(ts.todos[:index], ts.todos[index+1:]...)
	return nil
}

func (todos *Todos) Complete(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	ts := *todos
	if !ts.todos[index].Completed {
		ts.todos[index].Completed = true
		now := time.Now()
		ts.todos[index].CompletedAt = &now
	}

	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	if title == "" {
		return ErrEmptyTitle
	}

	ts := *todos
	ts.todos[index].Title = title

	return nil
}

func (todos *Todos) Render() error {
	return todos.render.Render(todos)
}

func (todos *Todos) Load() error {
	return todos.storage.Load(&(*todos).todos)
}

func (todos *Todos) Save() error {
	return todos.storage.Save((*todos).todos)
}
