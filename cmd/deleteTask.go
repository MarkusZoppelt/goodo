package cmd

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
)

func init() {
	rootCmd.AddCommand(deleteTaskCommand)
}

var deleteTaskCommand = &cobra.Command{
	Use:   "deleteTask [TODO INDEX] [TASK INDEX]",
	Short: "Delete a Task from a ToDo",
	Long:  `Delete a Task from a given Task`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		todos := db.GetAllTodos()
		if len(todos) == 0 {
			println("No ToDos found.")
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

		db.RemoveTaskFromTodo(selectedTodo, selectedTask)
	},
}
