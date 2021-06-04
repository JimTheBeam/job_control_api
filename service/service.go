package service

import (
	"job_control_api/config"
	"job_control_api/model"
	"job_control_api/repository"
)

// TaskService ...
type TaskService interface {
	CreateTask(*model.DBTask) error
}

// Service ...
type Service struct {
	Task TaskService
}

// NewService creates a new service
func NewService(repo *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		Task: NewTaskWebService(cfg, repo),
	}
}
