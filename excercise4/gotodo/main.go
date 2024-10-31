package main

import (
	"fmt"
	"gotodo/internal/gotodo"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
		os.Exit(1)
	}

	filename := os.Getenv("TODO_FILENAME")
	var storage gotodo.Storage
	switch {
	case filepath.Ext(filename) == ".json":
		storage = gotodo.NewJsonStorage(filename)
	case filepath.Ext(filename) == ".gob":
		storage = gotodo.NewGOBStorage(filename)
	default:
		fmt.Printf("Error: wrong filename or extension\n")
		os.Exit(1)
	}

	var render gotodo.Render
	command := gotodo.NewCmdFlags()
	switch {
	case command.List == gotodo.TableList:
		render = gotodo.NewTableRender(os.Stdout)

	case command.List == gotodo.CSVList:
		render = gotodo.NewCSVRender(os.Stdout)
	}
	todos := gotodo.NewTodos(render, storage)

	if err := todos.Load(); err != nil {
		fmt.Printf("Error Load: %s\n", err)
		os.Exit(1)
	}

	command.Run(todos)

	if err := todos.Save(); err != nil {
		fmt.Printf("Error Save: %s\n", err)
		os.Exit(1)
	}
}
