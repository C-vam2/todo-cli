package csv

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"example.com/todo-cli/internal/task"
	"github.com/gofrs/flock"
)

type CSVStorage struct {
	filePath string
}

func NewCSVStorage(filePath string) (*CSVStorage, error) {
	if filePath == "" {
		err := errors.New("Invalid file path.")
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	return &CSVStorage{filePath: filePath}, nil
}

func (storage *CSVStorage) AddTasks(ctx context.Context, t *task.Task) error {
	filepath := storage.filePath

	tasks, err := LoadTasks(filepath)

	if err != nil {
		return err
	}

	var maxId int = 0
	for _, value := range tasks {
		maxId = max(maxId, value.ID)
	}
	t.ID = maxId + 1
	tasks = append(tasks, *t)

	if err := SaveTasks(filepath, tasks); err != nil {
		return err
	}

	return nil

}

func (storage *CSVStorage) DeleteTask(ctx context.Context, id int) error {
	filePath := storage.filePath

	tasks, err := LoadTasks(filePath)

	if err != nil {
		return err
	}

	newTasks := []task.Task{}
	var taskFound bool
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		} else {
			taskFound = true
		}
	}

	if !taskFound {
		return fmt.Errorf("task with ID: %d not found", id)
	}

	if err := SaveTasks(filePath, newTasks); err != nil {
		return err
	}
	return nil
}

func (storage *CSVStorage) UpdateTask(ctx context.Context, t *task.Task) error {
	filePath := storage.filePath

	tasks, err := LoadTasks(filePath)

	if err != nil {
		return err
	}

	newTasks := []task.Task{}
	var taskFound bool
	for _, task := range tasks {
		if task.ID != t.ID {
			newTasks = append(newTasks, task)
		} else {
			if task.IsComplete {
				return fmt.Errorf("Task is already completed")
			}
			task.IsComplete = true
			newTasks = append(newTasks, task)
			taskFound = true
		}
	}

	if !taskFound {
		return fmt.Errorf("task with ID: %d not found", t.ID)
	}

	if err := SaveTasks(filePath, newTasks); err != nil {
		return err
	}
	return nil
}

func (storage *CSVStorage) GetTasks(ctx context.Context) ([]task.Task, error) {
	filepath := storage.filePath

	tasks, err := LoadTasks(filepath)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (storage *CSVStorage) GetTask(ctx context.Context, id int) (task.Task, error) {
	filePath := storage.filePath

	tasks, err := LoadTasks(filePath)
	if err != nil {
		return task.Task{}, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return task.Task{}, fmt.Errorf("task with ID: %d not found", id)

}

func SaveTasks(filepath string, tasks []task.Task) error {
	lock := flock.New(filepath + ".lock")
	err := lock.Lock()
	if err != nil {
		return err
	}
	defer lock.Unlock()

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	err = writer.Write([]string{"ID", "Description", "CreatedAt", "IsComplete"})

	if err != nil {
		return err
	}

	for _, task := range tasks {
		err = writer.Write([]string{strconv.Itoa(task.ID), task.Description, task.CreatedAt.Format(time.RFC3339), strconv.FormatBool(task.IsComplete)})
		if err != nil {
			return err
		}
	}
	writer.Flush()

	err = writer.Error()
	return err
}

func LoadTasks(filePath string) ([]task.Task, error) {

	lock := flock.New(filePath + ".lock")

	err := lock.Lock()

	if err != nil {
		return nil, err
	}

	defer lock.Unlock()

	file, err := os.Open(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return []task.Task{}, nil
		}
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	tasks := []task.Task{}

	for idx, row := range data {
		if idx == 0 {
			continue
		}

		if len(row) != 4 {
			continue
		}

		id, err := strconv.Atoi(row[0])

		if err != nil {
			// fmt.Println(err)
			continue
		}

		description := row[1]
		createdAt, err := time.Parse(time.RFC3339, row[2])

		if err != nil {
			// fmt.Println(err)
			continue
		}

		isComplete, err := strconv.ParseBool(row[3])

		if err != nil {
			// fmt.Println(err)
			continue
		}

		tasks = append(tasks, task.Task{ID: id, Description: description, CreatedAt: createdAt, IsComplete: isComplete})
	}

	return tasks, nil
}
