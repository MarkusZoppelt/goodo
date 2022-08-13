package db

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"zoppelt.net/goodo/task"
	"zoppelt.net/goodo/todo"
)

const dbFile string = "goodo.db"

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func CreateDB() {
	db, err := sql.Open("sqlite3", dbFile)
	check(err)
	defer db.Close()

	init := `
		DROP TABLE IF EXISTS todos;
		DROP TABLE IF EXISTS tasks;
		CREATE TABLE todos(id TEXT PRIMARY KEY, name TEXT, description TEXT, tasks BLOB);
		CREATE TABLE tasks(id TEXT PRIMARY KEY, name TEXT);
	`
	_, err = db.Exec(init)
	check(err)
}

func InsertTodo(t todo.Todo) {
	db, err := sql.Open("sqlite3", dbFile)
	check(err)
	defer db.Close()

	tasks, err := json.Marshal(t.Tasks)
	check(err)

	command := `
		INSERT INTO todos(id, name, description, tasks)
		VALUES(?,?,?,?)
	`
	_, err = db.Exec(command, t.ID, t.Name, t.Description, string(tasks))
	check(err)
}

func RemoveTodo(t todo.Todo) {
	db, err := sql.Open("sqlite3", dbFile)
	check(err)
	defer db.Close()

	_, err = db.Exec(`
		DELETE FROM todos
		WHERE id == ?
	`, t.ID)
	check(err)
}

func UpdateTodoTasks(todo todo.Todo, withTasks []task.Task) {
	db, err := sql.Open("sqlite3", dbFile)
	check(err)
	defer db.Close()

	tasks, err := json.Marshal(withTasks)
	check(err)

	_, err = db.Exec(`
		UPDATE todos
		SET tasks = ?
		WHERE id == ?
	`, string(tasks), todo.ID)
	check(err)
}

func AddTaskToTodo(to todo.Todo, ta task.Task) {
	todo := GetTodoWithID(to.ID)
	tasks := GetTasksForTodo(todo)
	tasks = append(tasks, ta)
	UpdateTodoTasks(todo, tasks)
}

func RemoveTaskFromTodo(to todo.Todo, ta task.Task) {
	todo := GetTodoWithID(to.ID)
	tasks := GetTasksForTodo(todo)
	var updatedTasks []task.Task
	for _, elem := range tasks {
		if elem.ID != ta.ID {
			updatedTasks = append(updatedTasks, elem)
		}
	}
	UpdateTodoTasks(todo, updatedTasks)
	removeTaskWithID(ta.ID)
}

func removeTaskWithID(tid string) {
	db, err := sql.Open("sqlite3", dbFile)
	check(err)
	defer db.Close()

	_, err = db.Exec(`
		DELETE FROM tasks
		WHERE id == ?
	`, tid)
	check(err)
}

func GetTasksForTodo(t todo.Todo) []task.Task {
	return t.Tasks
}

func GetTodoWithID(tid string) todo.Todo {
	db, err := sql.Open("sqlite3", dbFile)
	check(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM todos WHERE id == ?", tid)
	check(err)
	defer rows.Close()
	rows.Next()

	var id string
	var name string
	var description string
	var tasks_json string
	var tasks []task.Task

	err = rows.Scan(&id, &name, &description, &tasks_json)
	check(err)

	json.Unmarshal([]byte(tasks_json), &tasks)

	return todo.Todo{
		ID:          id,
		Name:        name,
		Description: description,
		Tasks:       tasks,
	}

}

func GetAllTodos() []todo.Todo {

	var todos []todo.Todo

	db, err := sql.Open("sqlite3", dbFile)
	check(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM todos")
	check(err)
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		var description string
		var tasks_json string
		var tasks []task.Task

		err = rows.Scan(&id, &name, &description, &tasks_json)
		check(err)

		err := json.Unmarshal([]byte(tasks_json), &tasks)
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
