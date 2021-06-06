package service

import (
	"job_control_api/model"
)

// CreateSubTask creates a new subtask
func (s *TaskWebService) CreateSubTask(subTask *model.DBSubTask) (model.DBSubTask, error) {

	// check if task exists
	_, err := s.repo.Repo.GetTask(subTask.TaskName)
	if err != nil {
		s.log.Warnf("service CreateSubTask get task: %v", err)
		return model.DBSubTask{}, err
	}

	// create a subtask
	err = s.repo.Repo.CreateSubTask(subTask)
	if err != nil {
		s.log.Errorf("service CreateSubTask: %v", err)
		return model.DBSubTask{}, err
	}

	// get a new created subtask
	newSubTask, err := s.repo.Repo.GetSubTask(subTask.Name)
	if err != nil {
		s.log.Errorf("service CreateSubTask get: %v", err)
		return model.DBSubTask{}, err
	}

	return newSubTask, nil
}

// GetSubTask get a subtask with name
func (s *TaskWebService) GetSubTask(name string) (model.DBSubTask, error) {

	subTask, err := s.repo.Repo.GetSubTask(name)
	if err != nil {
		s.log.Warnf("service GetTask: %v", err)
		return model.DBSubTask{}, err
	}
	return subTask, nil
}

// DeleteSubTask deletes a subtask with name
func (s *TaskWebService) DeleteSubTask(task *model.DBSubTask) error {

	err := s.repo.Repo.DeleteSubTask(task.Name)
	if err != nil {
		s.log.Errorf("service DeleteSubTask: %v", err)
		return err
	}

	return nil
}
