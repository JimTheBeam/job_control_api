package repository

import (
	"database/sql"
	"job_control_api/model"
	"job_control_api/repository/pg"
)

type TaskRepository interface {
	CreateTask(*model.DBTask) error
	GetTask(string) (model.DBTask, error)
	DeleteTask(string) error
}

// Repository ...
type Repository struct {
	Repo TaskRepository
}

// NewRepository ...
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Repo: pg.NewTaskPG(db),
	}
}
