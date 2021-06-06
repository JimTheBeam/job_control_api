package pg

import (
	"fmt"
	"job_control_api/model"
)

// CreateCost creates a new subtask
func (r *TaskPG) CreateCost(cost *model.DBCost) error {
	log := r.log
	log.Debug("DB: CreateCost start")
	defer log.Debug("DB: CreateCost end")

	sql := fmt.Sprintf("INSERT INTO costs (costs, subtask_name) VALUES ($1, $2)")

	if err := r.db.QueryRow(sql, cost.Cost, cost.SubTaskName).Err(); err != nil {
		log.Errorf("DB: CreateCost query: %v", err)
		return err
	}

	return nil
}

// GetCost gets subtask
func (r *TaskPG) GetCost(subTaskName string) (model.DBCost, error) {
	log := r.log
	log.Debug("DB: GetCost start")
	defer log.Debug("DB: GetCost end")

	var cost model.DBCost

	sql := fmt.Sprintf("SELECT costs, subtask_name FROM costs WHERE subtask_name=$1")

	err := r.db.QueryRow(sql, subTaskName).Scan(&cost.Cost, &cost.SubTaskName)
	if err != nil {
		log.Warnf("DB: GetCost: %v", err)
		return model.DBCost{}, err
	}

	return cost, nil
}

// DeleteCost deletes subtask with name
func (r *TaskPG) DeleteCost(subTaskName string) error {
	log := r.log
	log.Debug("DB: DeleteCost start")
	defer log.Debug("DB: DeleteCost end")

	sql := fmt.Sprintf("DELETE FROM costs WHERE subtask_name=$1")

	_, err := r.db.Exec(sql, subTaskName)
	if err != nil {
		log.Errorf("DB: DeleteCost: %v", err)
		return err
	}

	log.Debug("DB: Cost with subtask_name=%s deleted", subTaskName)

	return nil
}

// GetAllCost gets all costs for Cash
func (r *TaskPG) GetAllCost(cash *model.Data) error {
	log := r.log
	log.Debug("DB: GetAllCost start")
	defer log.Debug("DB: GetAllCost end")

	rows, err := r.db.Query("SELECT costs, subtask_name FROM costs")
	if err != nil {
		log.Errorf("DB:GetAllCost query: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		cost := model.SubTaskCost{}
		if err := rows.Scan(&cost.Costs, &cost.SubTaskName); err != nil {
			log.Errorf("DB:GetAllCost rows scan: %v", err)
			return err
		}
		// add cost to cash
		cash.Cost[cost.SubTaskName] = cost
	}
	if err := rows.Err(); err != nil {
		log.Errorf("DB:GetAllCost rows err: %v", err)
		return err
	}

	return nil
}
