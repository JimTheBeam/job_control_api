package service

import (
	"errors"
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
	log.Debug("CreateTask service starts")
	defer log.Debug("CreateTask service end")

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
	log := s.log
	log.Debug("GetTask service starts")
	defer log.Debug("GetTask service end")

	task, err := s.repo.Repo.GetTask(name)
	if err != nil {
		log.Errorf("service GetTask: %v", err)
		return model.DBTask{}, err
	}
	return task, nil
}

// DeleteTask deletes a task with name
func (s *TaskWebService) DeleteTask(task *model.DBTask) error {
	log := s.log
	log.Debug("DeleteTask service starts")
	defer log.Debug("DeleteTask service end")

	// check if task exists
	_, err := s.repo.Repo.GetTask(task.Name)
	if err != nil {
		log.Warningf("service DeleteTask task doesnt exist: %v", err)
		return errors.New("task doesnot exist")
	}
	// delete task from db
	err = s.repo.Repo.DeleteTask(task.Name)
	if err != nil {
		log.Errorf("service DeleteTask: %v", err)
		return errors.New("couldnot delete the task")
	}

	return nil
}

func (s *TaskWebService) GetAllTasks() {

}
