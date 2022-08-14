package cmd

import (
	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
	"zoppelt.net/goodo/task"
)

func init() {
	rootCmd.AddCommand(createTaskCommand)
}

var createTaskCommand = &cobra.Command{
	Use:   "addTask [TODO UUID] [NAME]",
	Short: "Add a new Task to a given Todo",
	Long:  `Add a new Task with a given Name to a given Todo UUID`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		to := db.GetTodoWithID(args[0])
		ta := task.New(args[1])
		db.AddTaskToTodo(to, *ta)
		println("Created new task with ID: ", ta.ID)
	},
}
