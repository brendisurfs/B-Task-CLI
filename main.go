package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/brendisurfs/go-cli/cmd"
	"github.com/brendisurfs/go-cli/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	fmt.Println("db for tasks was created!")
	// Execute app
	must(cmd.RootCmd.Execute())
}

// small handler for our use case to catch errors.
func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
