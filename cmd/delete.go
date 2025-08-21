package cmd

import (
	"fmt"
	"strconv"

	"github.com/kuipid01/task-cli/internal/models"
	"github.com/kuipid01/task-cli/internal/storage"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a task ID")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		tasks, _ := storage.LoadTasks()
		newTasks := []models.Task{}
		found := false

		for _, t := range tasks {
			if t.ID != id {
				newTasks = append(newTasks, t)
			} else {
				found = true
				fmt.Println("Task deleted:", t.Description)
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
	rootCmd.AddCommand(deleteCmd)
}
