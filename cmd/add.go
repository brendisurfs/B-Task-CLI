package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/TwinProduction/go-color"
	"github.com/brendisurfs/b-task/db"
	"github.com/spf13/cobra"
)

var strArgs []string

var addCmd = cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",

	Run: func(cmd *cobra.Command, args []string) {
		strArgs = append(strArgs, args...)
		task := strings.Join(strArgs, " ")

		// db methods
		_, err := db.CreateTask(task)
		if err != nil {
			log.Fatal(color.Red+"task creation failed! damn. : -> "+color.Reset, err.Error())
			return
		}
		// print out the todo back to the user.
		fmt.Printf(color.Green+"added \"%s\" to your task list\n"+color.Reset, task)
	},
}

// WHAT IS AN INIT FUNC?

/*INIT is a function that can be run before main, almost like a precursor, rather than just having stuff inside the main .
this is good for setting stuff up, accessing env variables, other stuff.
Its good to know, just USE IT SPARINGLY.
It is bloaty and should only be used in a case like this.
save your package.

*/
func init() {
	RootCmd.AddCommand(&addCmd)
}
