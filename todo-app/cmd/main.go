package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/Neksys2303/PFMS/todo-app/comands"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add new task")

	flag.Parse()

	todos := &todo.Todo{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Sample todo")
		err := todos.Storage(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "invalid comand")
		os.Exit(0)
	}

}
