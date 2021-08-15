package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/TwinProduction/go-color"
	"github.com/brendisurfs/b-task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete/doing!",

	Run: func(cmd *cobra.Command, args []string) {

		var ids []int

		// parse the ids
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println(color.Red+"error: failed to parse argument:"+color.Reset, arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.GetAllTasks()
		if err != nil {
			log.Fatal(color.Red+"something went wrong with getting all tasks."+color.Reset, err)
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println(color.Yellow+"invalid task number:"+color.Reset, id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf(color.Red+"failed to mark \"%d\" as complete. Error: %s\n"+color.Reset, id, err)
			} else {
				fmt.Printf(color.Green+"Marked Item \"%d\" -> complete.\n"+color.Reset, id)
			}
		}
	},
}

// adds our do command.
func init() {
	RootCmd.AddCommand(doCmd)
}
