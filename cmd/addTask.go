package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
	"zoppelt.net/goodo/task"
)

func init() {
	rootCmd.AddCommand(addTaskCommand)
}

var addTaskCommand = &cobra.Command{
	Use:   "addTask [TODO INDEX] [NAME]",
	Short: "Add a new Task to a given ToDo",
	Long:  `Add a new Task with a given Name to a given ToDo UUID`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		todos := db.GetAllTodos()

		index, _ := strconv.Atoi(args[0])
		selected := todos[index-1]

		ta := task.New(args[1])
		db.AddTaskToTodo(selected, *ta)
		println("Created new task with ID: ", ta.ID)
	},
}
