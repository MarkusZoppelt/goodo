package cmd

import (
	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
)

func init() {
	rootCmd.AddCommand(deleteTaskCommand)
}

var deleteTaskCommand = &cobra.Command{
	Use:   "deleteTask [TODO UUID] [TASK UUID]",
	Short: "Delete a Task from a ToDo",
	Long:  `Delete a Task from a given Task`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		to := db.GetTodoWithID(args[0])

		for _, task := range to.Tasks {
			if task.ID == args[1] {
				db.RemoveTaskFromTodo(to, task)
				println("Deleted Task with ID: %s from ToDo with ID: ", task.ID, to.ID)
			}
		}
	},
}
