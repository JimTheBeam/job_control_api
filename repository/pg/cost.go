package pg

import (
	"fmt"
	"job_control_api/model"
	"log"
)

// CreateCost creates a new subtask
func (r *TaskPG) CreateCost(cost *model.DBCost) error {
	log.Printf("DB: CreateCost start")
	defer log.Printf("DB: CreateCost end")

	sql := fmt.Sprintf("INSERT INTO costs (costs, subtask_name) VALUES ($1, $2)")

	if err := r.db.QueryRow(sql, cost.Cost, cost.SubTaskName).Err(); err != nil {
		log.Printf("DB: CreateCost query: %v", err)
		return err
	}

	return nil
}

// GetCost gets subtask
func (r *TaskPG) GetCost(subTaskName string) (model.DBCost, error) {
	log.Printf("DB: GetCost start")
	defer log.Printf("DB: GetCost end")

	var cost model.DBCost

	sql := fmt.Sprintf("SELECT costs, subtask_name FROM costs WHERE subtask_name=$1")

	err := r.db.QueryRow(sql, subTaskName).Scan(&cost.Cost, &cost.SubTaskName)
	if err != nil {
		log.Printf("DB: GetCost: %v", err)
		return model.DBCost{}, err
	}

	return cost, nil
}

// DeleteCost deletes subtask with name
func (r *TaskPG) DeleteCost(subTaskName string) error {
	log.Printf("DB: DeleteCost start")
	defer log.Printf("DB: DeleteCost end")

	sql := fmt.Sprintf("DELETE FROM costs WHERE subtask_name=$1")

	_, err := r.db.Exec(sql, subTaskName)
	if err != nil {
		log.Printf("DB: DeleteCost: %v", err)
		return err
	}

	log.Printf("DB: Cost with subtask_name=%s deleted", subTaskName)

	return nil
}

// GetAllCost gets all costs for Cash
func (r *TaskPG) GetAllCost(cash *model.Data) error {
	log.Printf("DB: GetAllCost start")
	defer log.Printf("DB: GetAllCost end")

	rows, err := r.db.Query("SELECT costs, subtask_name FROM costs")
	if err != nil {
		log.Printf("DB: Get all costs: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		cost := model.SubTaskCost{}
		if err := rows.Scan(&cost.Costs, &cost.SubTaskName); err != nil {
			return err
		}
		// add cost to cash
		cash.Cost[cost.SubTaskName] = cost
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
