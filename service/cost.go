package service

import (
	"errors"
	"job_control_api/model"
)

// CreateCost creates a new cost for subtask
func (s *TaskWebService) CreateCost(cost *model.DBCost) (model.DBCost, error) {
	log := s.log
	log.Debug("CreateCost service starts")
	defer log.Debug("CreateCost service end")

	// check if cost with name already exist
	_, err := s.repo.Repo.GetCost(cost.Name)
	if err == nil {
		log.Warnf("service CreateCost get cost: %v", err)
		return model.DBCost{}, errors.New("cost already exists")
	}

	// check if subtask exists
	_, err = s.repo.Repo.GetSubTask(cost.SubTaskName)
	if err != nil {
		log.Warnf("service CreateCost get subtask: %v", err)
		return model.DBCost{}, errors.New("subtask doesnot exist")
	}

	// create a cost
	err = s.repo.Repo.CreateCost(cost)
	if err != nil {
		log.Errorf("service CreateCost: %v", err)
		return model.DBCost{}, errors.New("couldnot create cost")
	}

	// get a new created cost
	newCost, err := s.repo.Repo.GetCost(cost.Name)
	if err != nil {
		log.Errorf("service CreateCost get: %v", err)
		return model.DBCost{}, errors.New("couldnot create cost")
	}

	return newCost, nil
}

// GetCost get a cost with name
func (s *TaskWebService) GetCost(subTaskName string) (model.DBCost, error) {
	log := s.log
	log.Debug("GetCost service starts")
	defer log.Debug("GetCost service end")

	cost, err := s.repo.Repo.GetCost(subTaskName)
	if err != nil {
		log.Warnf("service GetCost: %v", err)
		return model.DBCost{}, err
	}
	return cost, nil
}

// DeleteCost deletes a cost with name
func (s *TaskWebService) DeleteCost(cost *model.DBCost) error {
	log := s.log
	log.Debug("DeleteCost service starts")
	defer log.Debug("DeleteCost service end")

	log.Debugf("service DeleteCost deleting '%v'", cost)

	// check if cost exists
	_, err := s.repo.Repo.GetCost(cost.Name)
	if err != nil {
		log.Warningf("service DeleteCost get cost: %v", err)
		return errors.New("cost doesnot exist")
	}
	// delete cost
	err = s.repo.Repo.DeleteCost(cost.Name)
	if err != nil {
		log.Errorf("service DeleteCost: %v", err)
		return errors.New("couldnot delete cost")
	}

	return nil
}
