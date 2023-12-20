package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"small_console_applications_go/interacting/todo"
	"strings"
)

// Hardcoding file name for now
var todoFileName = ".todo.json"

func main() {
	// Parsing command line flags
	add := flag.Bool("add", false, "Task to be included in the todo list")
	del := flag.Int("del", 0, "Item to be deleted")
	list := flag.Bool("list", false, "List all tasks")
	verbose := flag.Bool("verbose", false, "Verbose output")
	complete := flag.Int("complete", 0, "Item to be completed")
	hideComplete := flag.Bool("hide-complete", false, "Hide completed tasks")
	flag.Parse()
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}
	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Divide what to do based on the number of arguments provided
	switch {
	case *list:
		l.Print(*verbose, *hideComplete)
	case *complete > 0:
		// Complete a task from the list
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Save the list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		// when any arguments (exclusing flags) are provided they will be used as the new task
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)
		// Save the list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *del > 0:
		// Delete a task from the list
		if err := l.Delete(*del); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Save the list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flag provided
		fmt.Println("Invalid option")
		os.Exit(1)
	}
}

// getTask function decides where to get the description for a new task from: argumetns or STDIN
func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}
	return s.Text(), nil
}
