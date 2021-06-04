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
func (s *TaskWebService) CreateTask(task *model.DBTask) (model.DBTask, error) {

	if err := s.repo.Repo.CreateTask(task); err != nil {
		log.Printf("service CreateTask: %v", err)
		return model.DBTask{}, err
	}

	newTask, err := s.repo.Repo.GetTask(task.Name)
	if err != nil {
		log.Printf("service GetTask: %v", err)
		return model.DBTask{}, err
	}

	return newTask, nil
}

// GetTask creates a new task
func (s *TaskWebService) GetTask(name string) (model.DBTask, error) {

	task, err := s.repo.Repo.GetTask(name)
	if err != nil {
		log.Printf("service GetTask: %v", err)
		return model.DBTask{}, err
	}
	return task, nil
}

// DeleteTask deletes a task with name
func (s *TaskWebService) DeleteTask(task *model.DBTask) error {

	err := s.repo.Repo.DeleteTask(task.Name)
	if err != nil {
		log.Printf("servise DeleteTask: %v", err)
		return err
	}

	return nil
}
