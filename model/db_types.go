package model

// DBTask is a database task
type DBTask struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// DBSubTask is a database sub-task
type DBSubTask struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`

	TaskName string `json:"task_name" validate:"required"`
}

// DBCost is a database sub-task cost
type DBCost struct {
	Cost        string `json:"cost" validate:"required"`
	Description string `json:"description" validate:"required"`

	SubTaskName string `json:"subtask_name" validate:"required"`
}
