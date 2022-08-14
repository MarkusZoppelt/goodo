package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
)

func init() {
	rootCmd.AddCommand(updateTodoCommand)
}

var updateTodoCommand = &cobra.Command{
	Use:   "updateTodo [TODO INDEX] [NAME] [DESCRIPTION]",
	Short: "Update a ToDo's name and description",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		todos := db.GetAllTodos()

		index, _ := strconv.Atoi(args[0])
		selected := todos[index-1]

		db.UpdateTodoName(selected, args[1])
		db.UpdateTodoDescription(selected, args[2])
	},
}
