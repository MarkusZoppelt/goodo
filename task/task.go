package task

import "github.com/google/uuid"

type Task struct {
	ID   string
	Name string
}

func New(name string) *Task {
	t := Task{
		ID:   uuid.New().String(),
		Name: name,
	}
	return &t
}
