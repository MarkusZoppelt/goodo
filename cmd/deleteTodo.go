package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
)

func init() {
	rootCmd.AddCommand(deleteTodoCommand)
}

var deleteTodoCommand = &cobra.Command{
	Use:   "deleteTodo [TODO INDEX]",
	Short: "Delete a ToDo",
	Long:  `Delete a ToDo and all its Tasks`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todos := db.GetAllTodos()

		index, _ := strconv.Atoi(args[0])
		selected := todos[index-1]

		db.RemoveTodo(selected)
		println("Deleted todo with ID: ", selected.ID)
	},
}
