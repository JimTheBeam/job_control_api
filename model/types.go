package model

import "time"

// Data represents all data
type Data struct {
	Task    map[string]Task
	SubTask map[string]SubTask
	Cost    map[string]SubTaskCost
}

type NewData struct {
	Task map[string]NTask
}

type NTask struct {
	Name        string
	Description string

	AverageCost time.Duration
	TotalCost   time.Duration

	SubTasks map[string]NSubTask
}

type NSubTask struct {
	Name        string
	Description string
	TaskName    string

	AverageCost time.Duration
	TotalCost   time.Duration

	PSubTaskCost []*SubTaskCost
}

type NCost struct {
	SubTaskName string
	Costs       time.Duration
}

// Task represents a task
type Task struct {
	Name        string
	Description string

	AverageCost time.Duration
	TotalCost   time.Duration

	PSubTasks []*SubTask
}

// SubTask represents a subtask
type SubTask struct {
	PTask *Task

	Name        string
	Description string
	TaskName    string

	PSubTaskCost *SubTaskCost
}

// SubTaskCost represents time cost for a subtask
type SubTaskCost struct {
	PSubTask *SubTask

	SubTaskName string
	Costs       time.Duration
}
