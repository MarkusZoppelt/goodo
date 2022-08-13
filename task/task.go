package task

import "github.com/google/uuid"

type Task struct {
	id   uuid.UUID
	Name string
}

func New(name string) *Task {
	t := Task{
		id:   uuid.New(),
		Name: name,
	}
	return &t
}
