package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"example.com/todo-cli/internal/storage"
	"example.com/todo-cli/internal/task"
	"github.com/mergestat/timediff"
)

const dataFile = "tasks.csv"

func Execute() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo-cli <command>")
		return
	}

	switch os.Args[1] {

	case "list":
		handleList()

	case "add":
		handleAdd()

	case "complete":
		handleComplete()

	case "delete":
		handleDelete()

	default:
		fmt.Fprintln(os.Stderr, "Unknown command:", os.Args[1])
	}

}

func handleList() {
	tasks, err := storage.LoadTasks(dataFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	showAll := len(os.Args) >= 3 && os.Args[2] == "-a"
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

func handleAdd() {

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Please provide description for your TODO task\n Format: todo-cli add '<description>'")
		return
	}

	newTask := strings.Join(os.Args[2:], " ")

	if newTask == "" {
		fmt.Fprintln(os.Stderr, "Your task is empty")
		return
	}

	tasks, err := storage.LoadTasks(dataFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	newTasks := task.AddTask(tasks, newTask)
	err = storage.SaveTasks(dataFile, newTasks)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("Added task:", newTask)
}

func handleComplete() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Missing task ID. Required Format: todo-cli complete <ID>")
		return
	}

	taskId, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid taskId, %s\n", os.Args[2])
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

func handleDelete() {

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Missing task ID. Required format: todo-cli delete <taskID>\n")
		return
	}

	taskID, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid taskID, %s \n", os.Args[2])
		return
	}

	tasks, err := storage.LoadTasks(dataFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	tasks, err = task.DeleteTask(tasks, taskID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	if err := storage.SaveTasks(dataFile, tasks); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	fmt.Printf("Task with taskID : %d deleted successfully\n", taskID)

}
