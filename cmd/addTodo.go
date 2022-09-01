package cmd

import (
	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
	"zoppelt.net/goodo/todo"
)

func init() {
	rootCmd.AddCommand(addTodoCommand)
}

var addTodoCommand = &cobra.Command{
	Use:   "addTodo [NAME] [DESCRIPTION]",
	Short: "Add a new ToDo",
	Long:  `Add a new ToDo with a given name and description.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		t := todo.New(args[0], args[1])
		db.InsertTodo(*t)
	},
}
