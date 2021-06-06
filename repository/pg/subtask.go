package pg

import (
	"fmt"
	"job_control_api/model"
	"log"
)

// CreateSubTask creates a new subtask
func (r *TaskPG) CreateSubTask(subTask *model.DBSubTask) error {
	log.Printf("DB: CreateSubTask start")
	defer log.Printf("DB: CreateSubTask end")

	sql := fmt.Sprintf("INSERT INTO sub_tasks (name, description, task_name) VALUES ($1, $2, $3)")

	if err := r.db.QueryRow(sql, subTask.Name, subTask.Description, subTask.TaskName).Err(); err != nil {
		log.Printf("DB: CreateSubTask query: %v", err)
		return err
	}

	return nil
}

// GetSubTask gets subtask
func (r *TaskPG) GetSubTask(name string) (model.DBSubTask, error) {
	log.Printf("DB: GetSubTask start")
	defer log.Printf("DB: GetSubTask end")

	var subTask model.DBSubTask

	sql := fmt.Sprintf("SELECT name, description, task_name FROM sub_tasks WHERE name=$1")

	err := r.db.QueryRow(sql, name).Scan(&subTask.Name, &subTask.Description, &subTask.TaskName)
	if err != nil {
		log.Printf("DB: GetSubTask: %v", err)
		return model.DBSubTask{}, err
	}

	return subTask, nil
}

// DeleteSubTask deletes subtask with name
func (r *TaskPG) DeleteSubTask(name string) error {
	log.Printf("DB: DeleteSubTask start")
	defer log.Printf("DB: DeleteSubTask end")

	sql := fmt.Sprintf("DELETE FROM sub_tasks WHERE name=$1")

	_, err := r.db.Exec(sql, name)
	if err != nil {
		log.Printf("DB: DeleteSubTask: %v", err)
		return err
	}

	log.Printf("DB: SubTask with name=%s deleted", name)

	return nil
}

// GetAllSubTasks gets all subtasks for Cash
func (r *TaskPG) GetAllSubTasks(cash *model.Data) error {
	log.Printf("DB: GetAllSubTasks start")
	defer log.Printf("DB: GetAllSubTasks end")

	rows, err := r.db.Query("SELECT name, description, task_name FROM sub_tasks ORDER BY name")
	if err != nil {
		log.Printf("DB: Get all subtasks: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		subTask := model.SubTask{}
		if err := rows.Scan(&subTask.Name, &subTask.Description, &subTask.TaskName); err != nil {
			return err
		}
		// add subtask to cash
		cash.SubTask[subTask.Name] = subTask
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
