package task

type Task struct {
	Name string
}

func New(name string) *Task {
	t := Task{Name: name}
	return &t
}
