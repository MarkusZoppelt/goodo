package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
)

func init() {
	rootCmd.AddCommand(listCommand)
}

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Prints all ToDos and Tasks",
	Long:  `Shows a pretty-format for all ToDos and their associated Tasks.`,
	Run: func(cmd *cobra.Command, args []string) {

		colorGreen := "\033[32m"
		colorReset := "\033[0m"

		todos := db.GetAllTodos()
		for i, todo := range todos {
			fmt.Print(string(colorGreen), "[", i+1, "] ", string(colorReset))
			fmt.Println(todo.ToString())
		}
	},
}
