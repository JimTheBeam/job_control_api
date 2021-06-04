package service

import (
	"job_control_api/model"
	"log"
)

// CreateCost creates a new subtask
func (s *TaskWebService) CreateCost(cost *model.DBCost) (model.DBCost, error) {

	// check if subtask exists
	_, err := s.repo.Repo.GetSubTask(cost.SubTaskName)
	if err != nil {
		log.Printf("service CreateCost get task: %v", err)
		return model.DBCost{}, err
	}

	// create a cost
	err = s.repo.Repo.CreateCost(cost)
	if err != nil {
		log.Printf("service CreateCost: %v", err)
		return model.DBCost{}, err
	}

	// get a new created cost
	newCost, err := s.repo.Repo.GetCost(cost.SubTaskName)
	if err != nil {
		log.Printf("service CreateCost get: %v", err)
		return model.DBCost{}, err
	}

	return newCost, nil
}

// GetCost get a subtask with name
func (s *TaskWebService) GetCost(subTaskName string) (model.DBCost, error) {

	cost, err := s.repo.Repo.GetCost(subTaskName)
	if err != nil {
		log.Printf("service GetCost: %v", err)
		return model.DBCost{}, err
	}
	return cost, nil
}

// DeleteCost deletes a subtask with name
func (s *TaskWebService) DeleteCost(cost *model.DBCost) error {

	err := s.repo.Repo.DeleteCost(cost.SubTaskName)
	if err != nil {
		log.Printf("service GetCost: %v", err)
		return err
	}

	return nil
}
