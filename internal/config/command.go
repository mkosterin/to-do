package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"to-do/internal/todo"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new item in the to-do list")
	flag.StringVar(&cf.Edit, "edit", "", "Edit an item by index and new description. id:newDescription")
	flag.IntVar(&cf.Del, "del", -1, "Specify item by id to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify item by id to toggle complete true/false")
	flag.BoolVar(&cf.List, "list", false, "Print all items")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *todo.Todos) {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("error: invalid format for edit. Please use id:newDescription")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("error: invalid index for edit")
			os.Exit(1)
		}
		todos.Edit(index, parts[1])
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	case cf.Del != -1:
		todos.Delete(cf.Del)
	default:
		fmt.Println("invalid command")
	}
}
