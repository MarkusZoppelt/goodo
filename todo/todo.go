package todo

import "zoppelt.net/goodo/task"

type Todo struct {
	Name        string
	Description string
	Tasks       []task.Task
}

func New(name string, description string) *Todo {
	var tasks []task.Task
	t := Todo{Name: name, Description: description, Tasks: tasks}
	return &t
}
