package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"example.com/todo-cli/internal/task"
	"github.com/spf13/cobra"
)

func NewAddCmd(repo task.TaskRepository) *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {

			description := strings.Join(args, " ")

			if description == "" {
				fmt.Fprintln(os.Stderr, "Your task is empty")
				return
			}

			t := task.Task{
				Description: description,
				CreatedAt:   time.Now(),
				IsComplete:  false,
			}

			if err := repo.AddTasks(context.Background(), &t); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		},
		Args:  cobra.MinimumNArgs(1),
		Use:   "add [task]",
		Short: "add a new task",
	}
}
