package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	db "zoppelt.net/goodo/db"

	task "zoppelt.net/goodo/task"
	todo "zoppelt.net/goodo/todo"
)

func showTodoList() {
	todos := db.GetAllTodos()
	for _, todo := range todos {
		fmt.Println(todo.ToString())
	}
}

func main() {

	fresh := flag.Bool("f", false, "Delete DB and start fresh. Deletes everything!")
	// showAll := flag.Bool("l", false, "Shows your ToDo List. Prints everything.")

	flag.Parse()

	if _, err := os.Stat(db.DataBaseFile); errors.Is(err, os.ErrNotExist) || *fresh {
		db.InitDB()
	}

	test := todo.New("Test Todo", "Test description")
	testTask1 := task.New("123")
	testTask2 := task.New("456")
	db.InsertTodo(*test)
	db.AddTaskToTodo(*test, *testTask1)
	db.AddTaskToTodo(*test, *testTask2)

	showTodoList()

	db.RemoveTaskFromTodo(*test, *testTask1)
	db.RemoveTaskFromTodo(*test, *testTask2)

	println("---")
	showTodoList()

	testTask3 := task.New("task that needs to be deleted")
	db.AddTaskToTodo(*test, *testTask3)

	println("---")
	showTodoList()

	db.RemoveTaskFromTodo(*test, *testTask3)

	println("---")
	showTodoList()
}
