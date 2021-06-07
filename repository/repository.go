package repository

import (
	"database/sql"
	"job_control_api/model"
	"job_control_api/repository/pg"

	"github.com/sirupsen/logrus"
)

type TaskRepository interface {
	CreateTask(*model.DBTask) error
	GetTask(string) (model.DBTask, error)
	DeleteTask(string) error
	GetAllTasks(*model.Data) error

	CreateSubTask(*model.DBSubTask) error
	GetSubTask(string) (model.DBSubTask, error)
	DeleteSubTask(string) error
	UpdateSubTask(*model.DBSubTask) error
	GetAllSubTasks(*model.Data) error

	CreateCost(*model.DBCost) error
	GetCost(string) (model.DBCost, error)
	DeleteCost(string) error
	GetAllCost(*model.Data) error
}

// Repository ...
type Repository struct {
	Repo TaskRepository
}

// NewRepository ...
func NewRepository(db *sql.DB, log *logrus.Logger) *Repository {
	return &Repository{
		Repo: pg.NewTaskPG(db, log),
	}
}
