package cmd

import (
	"fmt"
	"log"
	"strconv"

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
				fmt.Println("error: failed to parse argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.GetAllTasks()
		if err != nil {
			log.Fatal("something went wrong with getting all tasks.", err)
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("invalid task number: ", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("failed to mark \"%d\" as complete. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as complete.", id)
			}
		}

		fmt.Println(ids)
	},
}

// adds our do command.
func init() {
	RootCmd.AddCommand(doCmd)
}
