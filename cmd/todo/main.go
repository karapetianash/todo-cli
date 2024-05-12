package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/karapetianash/todo-cli"
)

var todoFileName = ".todo.json"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for educational purposes.\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	flag.BoolVar(&todo.AddFlag, "add", false, "Task to be included in the ToDo list.")
	flag.BoolVar(&todo.ListFlag, "list", false, "List all tasks.")
	flag.IntVar(&todo.CompleteFlag, "complete", 0, "Item to be completed.")
	flag.IntVar(&todo.DelFlag, "del", 0, "Item to be deleted.")
	flag.BoolVar(&todo.VerboseFlag, "verbose", false, "List all tasks with verbose information.")
	flag.BoolVar(&todo.ActiveFlag, "active", false, "List only not completed tasks.")

	flag.Parse()

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case todo.ListFlag || todo.ActiveFlag || todo.VerboseFlag:
		fmt.Print(l)
	case todo.CompleteFlag > 0:
		if err := l.Complete(todo.CompleteFlag); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case todo.AddFlag:
		// When any arguments (excluding flags) are provided they will be used as a new task
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)

		if err = l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case todo.DelFlag > 0:
		if err := l.Delete(todo.DelFlag); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option.")
		os.Exit(1)
	}
}

// getTask decides where to get the description of a new task from: arguments or STDIN
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
		return "", fmt.Errorf("task cannot be blanked")
	}

	return s.Text(), nil
}
