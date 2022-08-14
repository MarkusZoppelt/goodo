package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
)

func init() {
	rootCmd.AddCommand(updateTaskCommand)
}

var updateTaskCommand = &cobra.Command{
	Use:   "updateTask [TODO INDEX] [TASK INDEX] [NAME]",
	Short: "Update a ToDo's Task with a new name",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		todos := db.GetAllTodos()

		todoIndex, _ := strconv.Atoi(args[0])
		selectedTodo := todos[todoIndex-1]

		taskIndex, _ := strconv.Atoi(args[1])
		selectedTask := selectedTodo.Tasks[taskIndex-1]

		db.UpdateTask(selectedTodo, selectedTask, args[2])
		println("Updated Task with new name: ", args[2])
	},
}
