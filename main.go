package main

import (
	"flag"
	"fmt"

	task "zoppelt.net/goodo/task"
	todo "zoppelt.net/goodo/todo"
)

func main() {

	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "Be verbose. Prints more information")
	flag.Parse()

	todo1 := todo.New("Wash feet", "Wash both feet")
	task1 := task.New("Wash left foot")
	task2 := task.New("Wash right foot")

	a := append(todo1.Tasks, *task1, *task2)

	fmt.Println("Todo1: ", todo1.Name, "Description: ", todo1.Description)
	fmt.Println("Tasks for Todo1: ", a)
}
