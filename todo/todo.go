package todo

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"zoppelt.net/goodo/task"
)

type Todo struct {
	id          uuid.UUID
	Name        string
	Description string
	Tasks       []task.Task
}

func New(name string, description string) *Todo {
	var tasks []task.Task
	t := Todo{
		id:          uuid.New(),
		Name:        name,
		Description: description,
		Tasks:       tasks,
	}
	return &t
}

func (t Todo) ToString() string {
	pretty := t.Name + " - " + t.Description + "\n"

	for _, task := range t.Tasks {
		pretty += fmt.Sprintln(" ↳ ", task.Name)
	}

	return strings.TrimSuffix(pretty, "\n")
}

func (t *Todo) AddTask(task task.Task) {
	t.Tasks = append(t.Tasks, task)
}
