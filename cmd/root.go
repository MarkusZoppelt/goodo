package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"zoppelt.net/goodo/db"
)

var rootCmd = &cobra.Command{
	Use:   "goodo",
	Short: "GooDo is a simple CLI app for managing ToDos and Tasks",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {

	var DELETE bool
	rootCmd.Flags().BoolVarP(&DELETE, "DELETE", "D", false, "Deletes everything in the database.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// init db if it doesn't exist
	if _, err := os.Stat(db.DataBaseFile); errors.Is(err, os.ErrNotExist) || DELETE {
		db.InitDB()
	}
}
