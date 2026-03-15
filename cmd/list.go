package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"example.com/todo-cli/internal/storage/csv"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		handleList(showAll)
	},
	Use:   "list",
	Short: "List pending tasks (use -a to include completed)",
}
var showAll bool

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "show completed tasks")
}

func handleList(showAll bool) {
	tasks, err := csv.LoadTasks(dataFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	if showAll {
		fmt.Fprintln(w, "ID\tDescription\tCreated\tDone")
	} else {
		fmt.Fprintln(w, "ID\tDescription\tCreated")
	}

	for _, record := range tasks {
		if showAll {
			fmt.Fprintf(w, "%d\t%s\t%s\t%t\n", record.ID, record.Description, timediff.TimeDiff(record.CreatedAt), record.IsComplete)
		} else if !record.IsComplete {
			fmt.Fprintf(w, "%d\t%s\t%s\n", record.ID, record.Description, timediff.TimeDiff(record.CreatedAt))
		}
	}
	w.Flush()
}
