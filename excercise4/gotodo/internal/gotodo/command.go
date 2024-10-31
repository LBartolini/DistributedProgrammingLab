package gotodo

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add       string
	Del       int
	Edit      string
	Completed int
	List      string
}

const (
	TableList = "Table"
	CSVList   = "CSV"
)

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo, give me the title")
	flag.IntVar(&cf.Del, "remove", -1, "Remove a todo, give me its id")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo, give me the index and a new title. id:new_title")
	flag.IntVar(&cf.Completed, "complete", -1, "Specify the index of the todo to mark as completed")
	flag.StringVar(&cf.List, "list", "", "List all the todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Run(todos *Todos) {
	switch {
	case cf.List != "":
		err := todos.Render()

		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)

		if len(parts) != 2 {
			fmt.Printf("Error: invalid format for edit. Use id:new_title")
			os.Exit(1)
		}

		id, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Printf("Error: invalid id for edit. Use id:new_title")
			os.Exit(1)
		}

		err = todos.Edit(id, parts[1])

		if err != nil {
			fmt.Printf("Error: %s. Use id:new_title", err)
			os.Exit(1)
		}
	case cf.Completed != -1:
		err := todos.Complete(cf.Completed)
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
	default:
		flag.Usage()
		os.Exit(1)
	}
}
