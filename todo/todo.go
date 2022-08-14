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

	// colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorCyan := "\033[36m"
	colorReset := "\033[0m"

	pretty := string(colorCyan) + t.Name + string(colorReset) +
		" - " + t.Description + "\n"

	for i, task := range t.Tasks {
		pretty += "    ↳ " + string(colorYellow) + "[" + fmt.Sprint(i+1) + "] " +
			string(colorReset) + task.Name + "\n"
	}

	return strings.TrimSuffix(pretty, "\n")
}
