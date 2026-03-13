package task

import (
	"fmt"
	"time"
)

type TaskRepository interface {
	GetTasks() ([]Task, error)
	AddTasks(description string) (Task, error)
	CompleteTask(ID int) (Task, error)
	DeleteTask(ID int) error
}

type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	IsComplete  bool
}

func AddTask(tasks []Task, description string) []Task {
	var maxId int = 0
	for _, value := range tasks {
		maxId = max(maxId, value.ID)
	}
	tasks = append(tasks, Task{ID: maxId + 1, Description: description, CreatedAt: time.Now(), IsComplete: false})
	return tasks
}

func CompleteTask(tasks []Task, id int) error {

	for idx, task := range tasks {
		if task.ID == id {
			if task.IsComplete {
				return fmt.Errorf("task is already completed")
			}
			tasks[idx].IsComplete = true
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func DeleteTask(tasks []Task, id int) ([]Task, error) {
	newTasks := []Task{}
	var taskFound bool
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		} else {
			taskFound = true
		}
	}

	if !taskFound {
		return nil, fmt.Errorf("task with ID: %d not found", id)
	}

	return newTasks, nil
}
