package cmd

import (
	"fmt"
	"os"
	"strconv"

	"example.com/todo-cli/internal/storage"
	"example.com/todo-cli/internal/task"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		handleComplete(args)
	},
	Args:  cobra.MinimumNArgs(1),
	Use:   "complete [taskID]",
	Short: "Mark a task as completed",
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

func handleComplete(args []string) {
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "Missing task ID. Required Format: todo-cli complete <ID>")
		return
	}

	taskId, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid taskId, %s\n", args[2])
		return
	}

	tasks, err := storage.LoadTasks(dataFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s retrieving tasks from storage.Try again.\n", err)
		return
	}

	if err := task.CompleteTask(tasks, taskId); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s occured while updating the task\n", err)
		return
	}

	if err := storage.SaveTasks(dataFile, tasks); err != nil {
		fmt.Fprintf(os.Stderr, "Error : %s occured while saving the tasks. Try again.\n", err)
		return
	}

	fmt.Printf("Your task with taskID: %d has been marked as completed.\n", taskId)

}
