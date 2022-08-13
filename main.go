package main

import (
	"fmt"

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
	// var verbose bool
	// flag.BoolVar(&verbose, "verbose", false, "Be verbose. Prints more information")
	// flag.Parse()

	db.CreateDB()

	test := todo.New("Test Todo", "Test description")
	testTask1 := task.New("123")
	testTask2 := task.New("456")
	test.AddTask(*testTask1)
	test.AddTask(*testTask2)

	db.InsertTodo(*test)

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
