package pg

import (
	"fmt"
	"job_control_api/model"
)

// CreateCost creates a new cost
func (r *TaskPG) CreateCost(cost *model.DBCost) error {
	log := r.log
	log.Debug("DB: CreateCost start")
	defer log.Debug("DB: CreateCost end")

	sql := fmt.Sprintf("INSERT INTO costs (name, costs, subtask_name) VALUES ($1, $2, $3)")

	if err := r.db.QueryRow(sql, cost.Name, cost.Cost, cost.SubTaskName).Err(); err != nil {
		log.Errorf("DB: CreateCost query: %v", err)
		return err
	}

	return nil
}

// GetCost gets cost with name
func (r *TaskPG) GetCost(costName string) (model.DBCost, error) {
	log := r.log
	log.Debug("DB: GetCost start")
	defer log.Debug("DB: GetCost end")

	log.Debugf("DB:GetCost getting cost '%s'", costName)

	var cost model.DBCost

	sql := fmt.Sprintf("SELECT name, costs, subtask_name FROM costs WHERE name=$1")

	err := r.db.QueryRow(sql, costName).Scan(&cost.Name, &cost.Cost, &cost.SubTaskName)
	if err != nil {
		log.Warnf("DB: GetCost: %v", err)
		return model.DBCost{}, err
	}

	return cost, nil
}

// DeleteCost deletes cost with name
func (r *TaskPG) DeleteCost(costName string) error {
	log := r.log
	log.Debug("DB: DeleteCost start")
	defer log.Debug("DB: DeleteCost end")

	log.Debugf("DB: DeleteCost deleting cost '%s'", costName)

	sql := fmt.Sprintf("DELETE FROM costs WHERE name=$1")

	_, err := r.db.Exec(sql, costName)
	if err != nil {
		log.Errorf("DB: DeleteCost: %v", err)
		return err
	}

	log.Debugf("DB: Cost with subtask_name=%s deleted", costName)

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
