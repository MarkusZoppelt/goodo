package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	db "zoppelt.net/goodo/db"
	"zoppelt.net/goodo/task"
	"zoppelt.net/goodo/todo"
)

func showTodoList() {
	todos := db.GetAllTodos()
	for _, todo := range todos {
		fmt.Println(todo.ToString())
	}
}

func main() {

	flag.Usage = func() {

		fmt.Fprintf(flag.CommandLine.Output(), "Subcommands:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  createTodo <name> <description>    creates new Todo\n")

		fmt.Fprintf(flag.CommandLine.Output(), "Flags:\n")
		flag.PrintDefaults()
	}

	fresh := flag.Bool("f", false, "Delete DB and start fresh. Deletes everything!")
	showAll := flag.Bool("l", false, "Shows your ToDo List. Prints everything.")

	flag.Parse()

	if _, err := os.Stat(db.DataBaseFile); errors.Is(err, os.ErrNotExist) || *fresh {
		db.InitDB()
	}

	if *showAll {
		showTodoList()
	}

	if flag.Arg(0) == "createTodo" {
		t := todo.New(flag.Arg(1), flag.Arg(2))
		db.InsertTodo(*t)
		println("Created new todo with ID: ", t.ID)
	}

	if flag.Arg(0) == "addTask" {
		to := db.GetTodoWithID(flag.Arg(1))
		ta := task.New(flag.Arg(2))
		db.AddTaskToTodo(to, *ta)
		println("Created new task with ID: ", ta.ID)
	}

	if flag.Arg(0) == "deleteTodo" {
		t := db.GetTodoWithID(flag.Arg(1))
		db.RemoveTodo(t)
		println("Deleted todo with ID: ", t.ID)
	}
}
