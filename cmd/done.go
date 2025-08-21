package cmd

import (
	"fmt"
	"strconv"

	"github.com/kuipid01/task-cli/internal/models"
	"github.com/kuipid01/task-cli/internal/storage"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark task as done by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a task ID")

		}
		//convert id
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid Task ID")
		}

		//get tasks
		tasks, _ := storage.LoadTasks()
		newTasks := []models.Task{}
		found := false

		for _, v := range tasks {
			if v.ID != id {
				newTasks = append(newTasks, v)
			} else {
				found = true
				newTasks = append(newTasks, models.Task{ID: v.ID, Description: v.Description, Done: true})
			}

		}
		if !found {
			fmt.Println("Task not found")
			return
		}
		storage.SaveTasks(newTasks)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
