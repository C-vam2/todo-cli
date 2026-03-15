package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"example.com/todo-cli/internal/task"
	"github.com/spf13/cobra"
)

func NewUpdateCmd(repo task.TaskRepository) *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {

			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid taskId, %s\n", args[0])
				return
			}

			ctx := context.Background()
			task, err := repo.GetTask(ctx, taskId)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s retrieving tasks from storage.Try again.\n", err)
				return
			}
			task.IsComplete = true

			if err := repo.UpdateTask(ctx, &task); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

		},
		Args:  cobra.MinimumNArgs(1),
		Use:   "update [taskID]",
		Short: "Mark a task as completed",
	}
}
