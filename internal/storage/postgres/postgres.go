package postgres

import (
	"context"

	"example.com/todo-cli/internal/task"
	"github.com/jackc/pgx/v5"
)

type PostgresStorage struct {
	conn *pgx.Conn
}

func NewPostgresStorage(ctx context.Context, connStr string) (*PostgresStorage, error) {
	conn, err := pgx.Connect(ctx, connStr)

	if err != nil {
		return nil, err
	}

	return &PostgresStorage{conn: conn}, nil
}

func (postgres *PostgresStorage) DeleteTask(ctx context.Context, id int) error {
	conn := postgres.conn

	queryString := `
	DELETE FROM tasks
	WHERE ID = $1
	`

	_, err := conn.Exec(ctx, queryString, id)

	if err != nil {
		return err
	}

	return nil

}

func (postgres *PostgresStorage) GetTasks(ctx context.Context) ([]task.Task, error) {
	conn := postgres.conn

	queryString := `
	SELECT id,description,created_at,is_complete 
	FROM tasks
	`

	rows, err := conn.Query(ctx, queryString)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []task.Task

	for rows.Next() {
		var t task.Task
		err := rows.Scan(&t.ID, &t.Description, &t.CreatedAt, &t.IsComplete)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil { // why are we doing this as we are already returning the error inside the above loop
		return nil, err
	}

	return tasks, nil
}

func (postgres *PostgresStorage) AddTasks(ctx context.Context, t *task.Task) error {
	conn := postgres.conn

	queryString := `
		INSERT INTO tasks (description,created_at,is_complete)
		VALUES ($1,$2,$3)
		RETURNING id
	`

	row := conn.QueryRow(ctx, queryString, t.Description, t.CreatedAt, t.IsComplete)

	if err := row.Scan(&t.ID); err != nil {
		return err
	}

	return nil

}

func (postgres *PostgresStorage) UpdateTask(ctx context.Context, t *task.Task) error {
	conn := postgres.conn

	queryString := `
	UPDATE tasks
	SET description=$1,created_at=$2,is_complete=$3
	WHERE ID = $4
  `
	_, err := conn.Exec(ctx, queryString, t.Description, t.CreatedAt, t.IsComplete, t.ID)

	if err != nil {
		return err
	}

	return nil
}
