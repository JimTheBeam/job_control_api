package service

import (
	"errors"
	"job_control_api/model"
)

// CreateSubTask creates a new subtask
func (s *TaskWebService) CreateSubTask(subTask *model.DBSubTask) (model.DBSubTask, error) {
	log := s.log
	log.Debug("CreateSubTask service starts")
	defer log.Debug("CreateSubTask service end")

	// check if task exists
	_, err := s.repo.Repo.GetTask(subTask.TaskName)
	if err != nil {
		log.Warnf("service CreateSubTask get task: %v", err)
		return model.DBSubTask{}, err
	}

	// create a subtask
	err = s.repo.Repo.CreateSubTask(subTask)
	if err != nil {
		log.Errorf("service CreateSubTask: %v", err)
		return model.DBSubTask{}, err
	}

	// get a new created subtask
	newSubTask, err := s.repo.Repo.GetSubTask(subTask.Name)
	if err != nil {
		log.Errorf("service CreateSubTask get: %v", err)
		return model.DBSubTask{}, err
	}

	return newSubTask, nil
}

// GetSubTask get a subtask with name
func (s *TaskWebService) GetSubTask(name string) (model.DBSubTask, error) {
	log := s.log
	log.Debug("GetSubTask service starts")
	defer log.Debug("GetSubTask service end")

	subTask, err := s.repo.Repo.GetSubTask(name)
	if err != nil {
		log.Warnf("service GetTask: %v", err)
		return model.DBSubTask{}, err
	}
	return subTask, nil
}

// DeleteSubTask deletes a subtask with name
func (s *TaskWebService) DeleteSubTask(subTask *model.DBSubTask) error {
	log := s.log
	log.Debug("DeleteSubTask service starts")
	defer log.Debug("DeleteSubTask service end")

	// check if subtask exists
	_, err := s.repo.Repo.GetSubTask(subTask.Name)
	if err != nil {
		log.Warningf("service DeleteSubTask subtask doesnt exist: %v", err)
		return errors.New("subtask doesnot exist")
	}
	// delete subtask
	err = s.repo.Repo.DeleteSubTask(subTask.Name)
	if err != nil {
		log.Errorf("service DeleteSubTask: %v", err)
		return errors.New("couldnot delete the subtask")
	}

	return nil
}
