package cmd

import (
	"os"
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

		db.RemoveTodo(selected)
		println("Deleted todo with ID: ", selected.ID)
	},
}
