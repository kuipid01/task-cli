package cmd

import (
	"fmt"

	"github.com/kuipid01/task-cli/internal/models"
	"github.com/kuipid01/task-cli/internal/storage"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Long:  "Add a new task to the task list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please Provide a task")
			return
		}
		tasks, _ := storage.LoadTasks()
		task := models.Task{
			ID:          len(tasks) + 1,
			Description: args[0],
			Done:        false,
		}

		storage.SaveTasks(append(tasks, task))
		fmt.Println("Task added: ", task.Description)
	},
}

func init() {

	rootCmd.AddCommand(addCmd)
}
