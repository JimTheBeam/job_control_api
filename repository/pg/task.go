package pg

import (
	"database/sql"
	"fmt"
	"job_control_api/model"
	"log"
)

// TaskPG is a postgres task
type TaskPG struct {
	db *sql.DB
}

// NewTaskPG new task
func NewTaskPG(db *sql.DB) *TaskPG {
	return &TaskPG{
		db: db,
	}
}

// CreateTask creates a new task
func (r *TaskPG) CreateTask(task *model.DBTask) error {
	log.Printf("DB: CreateTask start")
	defer log.Printf("DB: CreateTask end")

	sql := fmt.Sprintf("INSERT INTO tasks (name, description) VALUES ($1, $2)")

	if err := r.db.QueryRow(sql, task.Name, task.Description).Err(); err != nil {
		log.Printf("DB: CreateTask query: %v", err)
		return err
	}

	return nil
}

// GetTask gets task
func (r *TaskPG) GetTask(name string) (model.DBTask, error) {
	log.Printf("DB: GetTask start")
	defer log.Printf("DB: GetTask end")

	var task model.DBTask

	sql := fmt.Sprintf("SELECT name, description FROM tasks WHERE name=$1")

	err := r.db.QueryRow(sql, name).Scan(&task.Name, &task.Description)
	if err != nil {
		log.Printf("DB: GetTask: %v", err)
		return model.DBTask{}, err
	}

	return task, nil
}

// DeleteTask deletes task with name
func (r *TaskPG) DeleteTask(name string) error {
	log.Printf("DB: DeleteTask start")
	defer log.Printf("DB: DeleteTask end")

	sql := fmt.Sprintf("DELETE FROM tasks WHERE name=$1")

	_, err := r.db.Exec(sql, name)
	if err != nil {
		log.Printf("DB: DeleteTask: %v", err)
		return err
	}

	log.Printf("DB: Task with name=%s deleted", name)

	return nil
}

// GetAllTasks gets all tasks for Cash
func (r *TaskPG) GetAllTasks(cash *model.Data) error {
	log.Printf("DB: GetAllTasks start")
	defer log.Printf("DB: GetAllTasks end")

	rows, err := r.db.Query("SELECT name, description FROM tasks ORDER BY name")
	if err != nil {
		log.Printf("DB: Get all tasks: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		task := model.Task{}
		if err := rows.Scan(&task.Name, &task.Description); err != nil {
			return err
		}
		// add task to cash
		cash.Task[task.Name] = task
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
