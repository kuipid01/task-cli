package cmd

import (
	"fmt"

	"github.com/kuipid01/task-cli/internal/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long:  "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, _ := storage.LoadTasks()
		if len(tasks) == 0 {
			fmt.Println("No tasks found")
		}

		for _, t := range tasks {
			status := "❌"
			if t.Done {
				status = "✅"
			}
			fmt.Printf("[%d] %s %s\n", t.ID, status, t.Description)

		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
