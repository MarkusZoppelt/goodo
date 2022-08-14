package cmd

import (
	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
	"zoppelt.net/goodo/todo"
)

func init() {
	rootCmd.AddCommand(createTodoCommand)
}

var createTodoCommand = &cobra.Command{
	Use:   "createTodo [NAME] [DESCRIPTION]",
	Short: "Creates a new ToDo",
	Long:  `Creates a new ToDo with a given name and description.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		t := todo.New(args[0], args[1])
		db.InsertTodo(*t)
		println("Created new todo with ID: ", t.ID)
	},
}
