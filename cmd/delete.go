package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"example.com/todo-cli/internal/task"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(repo task.TaskRepository) *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			taskID, err := strconv.Atoi(args[2])

			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid taskID, %s \n", args[0])
				return
			}

			if err := repo.DeleteTask(context.Background(), taskID); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		},
		Args:  cobra.MinimumNArgs(1),
		Use:   "delete [taskID]",
		Short: "Delete a given task",
	}
}
