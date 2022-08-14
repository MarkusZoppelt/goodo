package todo

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"zoppelt.net/goodo/task"
)

type Todo struct {
	ID          string
	Name        string
	Description string
	Tasks       []task.Task
}

func New(name string, description string) *Todo {
	var tasks []task.Task
	t := Todo{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Tasks:       tasks,
	}
	return &t
}

func (t Todo) ToString() string {
	pretty := t.Name + " - " + t.Description + "\n"

	for i, task := range t.Tasks {
		pretty += fmt.Sprintln("       ↳ [", i+1, "]", task.Name)
	}

	return strings.TrimSuffix(pretty, "\n")
}
