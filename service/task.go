package service

import (
	"job_control_api/config"
	"job_control_api/model"
	"job_control_api/repository"

	"github.com/sirupsen/logrus"
)

// TaskWebService ..
type TaskWebService struct {
	repo *repository.Repository
	cfg  *config.Config
	log  *logrus.Logger
}

// NewTaskWebService creates a new task web service
func NewTaskWebService(cfg *config.Config, repo *repository.Repository, log *logrus.Logger) *TaskWebService {
	return &TaskWebService{
		repo: repo,
		cfg:  cfg,
		log:  log,
	}
}

// CreateTask creates a new task
func (s *TaskWebService) CreateTask(task *model.DBTask) (model.DBTask, error) {
	log := s.log
	// create task
	if err := s.repo.Repo.CreateTask(task); err != nil {
		log.Errorf("service CreateTask: %v", err)
		return model.DBTask{}, err
	}

	// get new task
	newTask, err := s.repo.Repo.GetTask(task.Name)
	if err != nil {
		log.Errorf("service GetTask: %v", err)
		return model.DBTask{}, err
	}

	return newTask, nil
}

// GetTask gets a task with name
func (s *TaskWebService) GetTask(name string) (model.DBTask, error) {

	task, err := s.repo.Repo.GetTask(name)
	if err != nil {
		s.log.Errorf("service GetTask: %v", err)
		return model.DBTask{}, err
	}
	return task, nil
}

// DeleteTask deletes a task with name
func (s *TaskWebService) DeleteTask(task *model.DBTask) error {

	err := s.repo.Repo.DeleteTask(task.Name)
	if err != nil {
		s.log.Errorf("servise DeleteTask: %v", err)
		return err
	}

	return nil
}

func (s *TaskWebService) GetAllTasks() {

}
