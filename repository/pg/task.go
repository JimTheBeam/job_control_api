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

func (r *TaskPG) CreateTask(task *model.DBTask) error {
	log.Printf("DB: CreateTask start")
	defer log.Printf("DB: CreateTask end")

	sql := fmt.Sprintf("INSERT INTO tasks (name) VALUES ($1)")

	if err := r.db.QueryRow(sql, task.Name).Err(); err != nil {
		log.Printf("DB: CreateTask query: %v", err)
		return err
	}

	return nil
}

func (r *TaskPG) GetTask(name string) (*model.DBTask, error) {
	log.Printf("DB: GetTask start")
	defer log.Printf("DB: GetTask end")

	var task model.DBTask

	sql := fmt.Sprintf("SELECT name FROM tasks WHERE name=$1")

	err := r.db.QueryRow(sql, name).Scan(&task.Name)
	if err != nil {
		log.Printf("DB: GetTask: %v", err)
		return &model.DBTask{}, err
	}

	return &task, nil
}
