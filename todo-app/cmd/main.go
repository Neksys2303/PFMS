package main

import (
	"flag"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add new task")

	flag.Parse()

	todos := &todo.Todo{}

}
