package main

import (
	"errors"
	"os"

	"zoppelt.net/goodo/cmd"
	"zoppelt.net/goodo/db"
)

func main() {

	// init db if it doesn't exist
	if _, err := os.Stat(db.DataBaseFile); errors.Is(err, os.ErrNotExist) {
		db.InitDB()
	}

	cmd.Execute()
}
