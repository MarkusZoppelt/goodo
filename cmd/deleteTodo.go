package cmd

import (
	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
)

func init() {
	rootCmd.AddCommand(deleteTodoCommand)
}

var deleteTodoCommand = &cobra.Command{
	Use:   "deleteTodo [TODO UUID]",
	Short: "Delete a ToDo",
	Long:  `Delete a ToDo and all its Tasks`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		t := db.GetTodoWithID(args[0])
		db.RemoveTodo(t)
		println("Deleted todo with ID: ", t.ID)
	},
}
