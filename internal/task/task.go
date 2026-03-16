package task

import (
	"context"
	"time"
)

type TaskRepository interface {
	GetTasks(ctx context.Context) ([]Task, error)
	AddTasks(ctx context.Context, t *Task) error
	UpdateTask(ctx context.Context, t *Task) error
	DeleteTask(ctx context.Context, id int) error
	GetTask(ctx context.Context, id int) (Task, error)
}

type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	IsComplete  bool
}
