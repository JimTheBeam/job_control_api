package model

import "time"

// Data represents all data
type Data struct {
	Task    map[string]Task
	SubTask map[string]SubTask
	Cost    map[string]SubTaskCost

	Trees map[string]TreeRoot // why? maybe map[string]Task is enough
}

// Task represents a task
type Task struct {
	Name string

	AverageCost time.Duration
	TotalCost   time.Duration

	SubTask []*SubTask
}

// SubTask represents a sub-task
type SubTask struct {
	Task *Task

	Name string

	SubTaskCost []*SubTaskCost
}

// SubTaskCost represents time cost for a sub-task
type SubTaskCost struct {
	SubTask *SubTask

	Name  string
	Costs time.Duration
}

// TreeRoot represents root for a tree Task-SubTask-SubTaskCost
type TreeRoot struct {
	Task *Task
}
