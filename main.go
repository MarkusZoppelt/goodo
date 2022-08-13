package main

import (
	"fmt"

	task "zoppelt.net/goodo/task"
	todo "zoppelt.net/goodo/todo"
)

func main() {
	// var verbose bool
	// flag.BoolVar(&verbose, "verbose", false, "Be verbose. Prints more information")
	// flag.Parse()

	t := todo.New("Wash feet", "Wash both feet")
	t.AddTask(*task.New("Wash left foot"))
	t.AddTask(*task.New("Wash right foot"))

	fmt.Println(t.ToString())
}
