package cmd

import (
	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
)

func init() {
	rootCmd.AddCommand(updateTodoCommand)
}

var updateTodoCommand = &cobra.Command{
	Use:   "updateTodo [TODO UUID] [NAME] [DESCRIPTION]",
	Short: "Update a Todo's name and description",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		t := db.GetTodoWithID(args[0])
		db.UpdateTodoName(t, args[1])
		db.UpdateTodoDescription(t, args[2])
	},
}
