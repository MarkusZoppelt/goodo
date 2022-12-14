package cmd

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
	"zoppelt.net/goodo/task"
)

func init() {
	rootCmd.AddCommand(addTaskCommand)
}

var addTaskCommand = &cobra.Command{
	Use:   "addTask [TODO INDEX] [NAME]",
	Short: "Add a new Task to a given ToDo",
	Long:  `Add a new Task with a given Name to a given ToDo UUID`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		todos := db.GetAllTodos()
		if len(todos) == 0 {
			println("No ToDos found. \nHint: You can only add Tasks to existing ToDos")
			os.Exit(1)
		}

		index, err := strconv.Atoi(args[0])
		if err != nil || index < 1 || index > len(todos) {
			println("Invalid index")
			os.Exit(1)
		}

		selected := todos[index-1]

		ta := task.New(args[1])
		db.AddTaskToTodo(selected, *ta)
	},
}
