package db

import (
	"testing"

	"zoppelt.net/goodo/todo"
)

func TestInsertTodo(t *testing.T) {

	InitDB()

	test := todo.New("Test Todo", "Test Description")
	println(test.ToString())
	InsertTodo(*test)

	comp := GetTodoWithID(test.ID)

	println(comp.ToString())

	if test.ID != comp.ID || test.Name != comp.Name ||
		test.Description != comp.Description {
		t.Errorf("got %s with ID %s, wanted %s with ID %s", test.ToString(), test.ID, comp.ToString(), comp.ID)
	}
}
