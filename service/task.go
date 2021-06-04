package service

import (
	"job_control_api/config"
	"job_control_api/model"
	"job_control_api/repository"
	"log"
)

// TaskWebService ..
type TaskWebService struct {
	repo *repository.Repository
	cfg  *config.Config
}

// NewTaskWebService creates a new task web service
func NewTaskWebService(cfg *config.Config, repo *repository.Repository) *TaskWebService {
	return &TaskWebService{
		repo: repo,
		cfg:  cfg,
	}
}

// CreateTask creates a new task
func (s *TaskWebService) CreateTask(task *model.DBTask) error {

	if err := s.repo.Repo.CreateTask(task); err != nil {
		log.Printf("service CreateTask: %v", err)
	}

	return nil
}
