package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"example.com/todo-cli/internal/task"
	"github.com/gofrs/flock"
)

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
