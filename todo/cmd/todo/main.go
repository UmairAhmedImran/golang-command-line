package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
	"UmairAhmedImran/todo"
)
var todoFileName = ".todo.json"


func main() {

	
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME") 
	}
	add := flag.Bool("add", false, "Add task to the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	deleteTask := flag.Int("delete", 0, "Item to be deleted") 
	todoList := flag.Bool("todo", false, "Item to do")
	completedList := flag.Bool("completedtodo", false, "items completed")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
		"%s tool. Developed for The Pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2020\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	flag.Parse()

	l := &todo.List{}

	// Use the Get method to read to do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
		case *list:
			fmt.Print(l)
		case *complete > 0:
			if err := l.Complete(*complete); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		case *todoList:
			todoList := l.TodoList()
			fmt.Print(todoList, "oklies\n what now huh")
		case *completedList:
			completedTaskList := l.CompletedList()
			fmt.Print(completedTaskList)
		case *add:
			t, err := getTask(os.Stdin, flag.Args()...)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			l.Add(t)
			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		case *deleteTask > 0:
			if err := l.Delete(*deleteTask); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		default:
			fmt.Fprintln(os.Stderr, "Invalid Option")
			os.Exit(1)
	}

}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", nil
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}
	return s.Text(), nil
}
