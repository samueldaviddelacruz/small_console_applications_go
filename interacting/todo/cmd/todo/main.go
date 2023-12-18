package main

import (
	"fmt"
	"os"
	"small_console_applications_go/interacting/todo"
)

// Hardcoding file name for now
const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
