package cmd

import (
	"os"
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
		if len(todos) == 0 {
			println("No ToDos found.")
			os.Exit(1)
		}

		index, err := strconv.Atoi(args[0])
		if err != nil || index < 1 || index > len(todos) {
			println("Invalid index")
			os.Exit(1)
		}
		selected := todos[index-1]

		db.UpdateTodoName(selected, args[1])
		db.UpdateTodoDescription(selected, args[2])
	},
}
