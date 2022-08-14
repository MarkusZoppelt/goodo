package cmd

import (
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

		todoIndex, _ := strconv.Atoi(args[0])
		selectedTodo := todos[todoIndex-1]

		taskIndex, _ := strconv.Atoi(args[1])
		selectedTask := selectedTodo.Tasks[taskIndex-1]

		db.RemoveTaskFromTodo(selectedTodo, selectedTask)
		println("Deleted Task with ID:", selectedTask.ID)
	},
}
