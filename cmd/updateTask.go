package cmd

import (
	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
)

func init() {
	rootCmd.AddCommand(updateTaskCommand)
}

var updateTaskCommand = &cobra.Command{
	Use:   "updateTask [TODO UUID] [TASK UUID] [NAME]",
	Short: "Update a Todo's Task with a new name",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		to := db.GetTodoWithID(args[0])

		for _, task := range to.Tasks {
			if task.ID == args[1] {
				db.UpdateTask(task, args[2])
				println("Updated Task with ID: %s with new name: ", task.Name)
			}
		}
	},
}
