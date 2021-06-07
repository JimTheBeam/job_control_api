package cash

import (
	"fmt"
	"job_control_api/model"
	"job_control_api/repository"

	"github.com/sirupsen/logrus"
)

// Init initialize cash
func Init(r *repository.Repository, log *logrus.Logger) error {
	taskMap := make(map[string]model.Task)
	subTaskMap := make(map[string]model.SubTask)
	costMap := make(map[string]model.SubTaskCost)
	Cash := model.Data{
		Task:    taskMap,
		SubTask: subTaskMap,
		Cost:    costMap,
	}

	log.Info(Cash)

	// load cash from database
	err := loadFromDB(&Cash, r)
	if err != nil {
		log.Errorf("load cash: %v", err)
		return err
	}
	log.Debug("cach loaded from db")

	// fill relations
	createRelations(&Cash)

	// calc average and total costs
	CalcAvTotal(&Cash)

	for n, val := range Cash.Task {
		fmt.Printf("TaskName: '%s', Pointers: %v\n", n, val)
	}

	return nil
}

// loadFromDB loads data from database
func loadFromDB(cash *model.Data, r *repository.Repository) error {
	// load tasks
	err := r.Repo.GetAllTasks(cash)
	if err != nil {
		logrus.Errorf("could not load task to cash")
		return err
	}

	// load subtasks
	err = r.Repo.GetAllSubTasks(cash)
	if err != nil {
		logrus.Errorf("could not load subtask to cash")
		return err
	}

	// load costs
	err = r.Repo.GetAllCost(cash)
	if err != nil {
		logrus.Errorf("could not load costs to cash")
		return err
	}

	return nil
}

// createRelations fill pointers in data
func createRelations(cash *model.Data) {
	// for name := range cash.SubTask {
	// fill task pointers
	// taskName := cash.SubTask[name].TaskName
	// cash.SubTask[name].PTask = &cash.Task[taskName]

	// fill subtask pointer in task
	// cash.Task[taskName].PSubTasks = append(cash.Task[taskName].PSubTasks, cash.SubTask[name])
	// }

	// for subTaskName := range cash.Cost {
	// fill subtask pointers
	// cash.SubTask[subTaskName].PSubTaskCost = cash.Cost[subTaskName]

	// fill subtask pointer in cost
	// cash.Cost[subTaskName].PSubTask = cash.SubTask[subTaskName]

}

// CalcAvTotal calculate average cost and total cost for task
func CalcAvTotal(cash *model.Data) {
	for name := range cash.Task {
		if len(cash.Task[name].PSubTasks) == 0 {
			continue
		}

		// var sum time.Duration
		// var n time.Duration

		// for i := 0; i < len(cash.Task[name].PSubTasks); i++ {
		// 	sum += cash.Task[name].PSubTasks[i].PSubTaskCost.Costs
		// 	n++
		// }
		// cash.Task[name].TotalCost = sum
		// cash.Task[name].AverageCost = sum / n
	}
}
