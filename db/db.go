package db

import (
	"database/sql"
	_ "embed"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"zoppelt.net/goodo/task"
	"zoppelt.net/goodo/todo"
)

const DataBaseFile string = "goodo.db"

var (
	//go:embed schema.sql
	schemaSQL string

	//go:embed insertTodo.sql
	insertTodoSQL string

	//go:embed removeTodo.sql
	removeTodoSQL string

	//go:embed updateTodoName.sql
	updateTodoNameSQL string

	//go:embed updateTodoDescription.sql
	updateTodoDescriptionSQL string

	//go:embed updateTodoTasks.sql
	updateTodoTasksSQL string

	//go:embed deleteTask.sql
	deleteTaskSQL string

	//go:embed selectTodo.sql
	selectTodoSQL string

	//go:embed selectAllTodos.sql
	selectAllTodosSQL string
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func InitDB() {
	db, err := sql.Open("sqlite3", DataBaseFile)
	check(err)
	defer db.Close()

	_, err = db.Exec(schemaSQL)
	check(err)
}

func InsertTodo(t todo.Todo) {
	db, err := sql.Open("sqlite3", DataBaseFile)
	check(err)
	defer db.Close()

	tasks, err := json.Marshal(t.Tasks)
	check(err)

	_, err = db.Exec(insertTodoSQL, t.ID, t.Name, t.Description, string(tasks))
	check(err)
}

func RemoveTodo(t todo.Todo) {
	db, err := sql.Open("sqlite3", DataBaseFile)
	check(err)
	defer db.Close()

	_, err = db.Exec(removeTodoSQL, t.ID)
	check(err)

	// also delete every task from that ToDo from the db
	for _, task := range t.Tasks {
		deleteTaskWithID(task.ID)
	}
}

func UpdateTodoName(todo todo.Todo, withName string) {
	db, err := sql.Open("sqlite3", DataBaseFile)
	check(err)
	defer db.Close()

	_, err = db.Exec(updateTodoNameSQL, withName, todo.ID)
	check(err)
}

func UpdateTodoDescription(todo todo.Todo, withDescription string) {
	db, err := sql.Open("sqlite3", DataBaseFile)
	check(err)
	defer db.Close()

	_, err = db.Exec(updateTodoDescriptionSQL, withDescription, todo.ID)
	check(err)
}

func UpdateTodoTasks(todo todo.Todo, withTasks []task.Task) {
	db, err := sql.Open("sqlite3", DataBaseFile)
	check(err)
	defer db.Close()

	tasks, err := json.Marshal(withTasks)
	check(err)

	_, err = db.Exec(updateTodoTasksSQL, string(tasks), todo.ID)
	check(err)
}

func AddTaskToTodo(to todo.Todo, ta task.Task) {
	todo := GetTodoWithID(to.ID)
	tasks := todo.Tasks
	tasks = append(tasks, ta)
	UpdateTodoTasks(todo, tasks)
}

func UpdateTask(to todo.Todo, ta task.Task, withName string) {
	db, err := sql.Open("sqlite3", DataBaseFile)
	check(err)
	defer db.Close()

	todo := GetTodoWithID(to.ID)

	var updatedTasks []task.Task
	for _, elem := range todo.Tasks {
		if elem.ID == ta.ID {
			elem.Name = withName
		}
		updatedTasks = append(updatedTasks, elem)
	}
	UpdateTodoTasks(todo, updatedTasks)
}

func RemoveTaskFromTodo(to todo.Todo, ta task.Task) {
	todo := GetTodoWithID(to.ID)
	tasks := todo.Tasks
	var updatedTasks []task.Task
	for _, elem := range tasks {
		if elem.ID != ta.ID {
			updatedTasks = append(updatedTasks, elem)
		}
	}
	UpdateTodoTasks(todo, updatedTasks)
	deleteTaskWithID(ta.ID)
}

func deleteTaskWithID(tid string) {
	db, err := sql.Open("sqlite3", DataBaseFile)
	check(err)
	defer db.Close()

	_, err = db.Exec(deleteTaskSQL, tid)
	check(err)
}

func GetTodoWithID(tid string) todo.Todo {
	db, err := sql.Open("sqlite3", DataBaseFile)
	check(err)
	defer db.Close()

	rows, err := db.Query(selectTodoSQL, tid)
	check(err)
	defer rows.Close()
	rows.Next()

	var id string
	var name string
	var description string
	var tasksJSON string
	var tasks []task.Task

	err = rows.Scan(&id, &name, &description, &tasksJSON)
	check(err)

	json.Unmarshal([]byte(tasksJSON), &tasks)

	return todo.Todo{
		ID:          id,
		Name:        name,
		Description: description,
		Tasks:       tasks,
	}
}

func GetAllTodos() []todo.Todo {
	var todos []todo.Todo

	db, err := sql.Open("sqlite3", DataBaseFile)
	check(err)
	defer db.Close()

	rows, err := db.Query(selectAllTodosSQL)
	check(err)
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		var description string
		var tasksJSON string
		var tasks []task.Task

		err = rows.Scan(&id, &name, &description, &tasksJSON)
		check(err)

		err := json.Unmarshal([]byte(tasksJSON), &tasks)
		check(err)

		t := todo.Todo{
			ID:          id,
			Name:        name,
			Description: description,
			Tasks:       tasks,
		}
		todos = append(todos, t)
	}
	return todos
}
