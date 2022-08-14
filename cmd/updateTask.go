package cmd

import (
	"os"
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
		if len(todos) == 0 {
			println("No ToDos found. \nHint: You can only update Tasks of existing ToDos")
			os.Exit(1)
		}

		todoIndex, errTo := strconv.Atoi(args[0])
		if errTo != nil || todoIndex < 1 || todoIndex > len(todos) {
			println("Invalid ToDo index")
			os.Exit(1)
		}
		selectedTodo := todos[todoIndex-1]

		taskIndex, errTa := strconv.Atoi(args[1])
		if errTa != nil || taskIndex < 1 || taskIndex > len(selectedTodo.Tasks) {
			println("Invalid Task index")
			os.Exit(1)
		}
		selectedTask := selectedTodo.Tasks[taskIndex-1]

		db.UpdateTask(selectedTodo, selectedTask, args[2])
		println("Updated Task with new name: ", args[2])
	},
}
