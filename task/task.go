package task

import "github.com/google/uuid"

type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func New(name string) *Task {
	t := Task{
		ID:   uuid.New().String(),
		Name: name,
	}
	return &t
}
