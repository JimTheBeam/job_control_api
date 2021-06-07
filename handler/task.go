package handler

import (
	"job_control_api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateTask creates new task
func (h *TaskHandler) CreateTask(ctx echo.Context) error {
	log := h.log
	log.Debug("CreateTask handler starts")
	defer log.Debug("CreateTask handler end")

	var task model.DBTask

	if err := ctx.Bind(&task); err != nil {
		log.Warningf("Bind json Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	// validate task
	if err := ctx.Validate(&task); err != nil {
		log.Warningf("Validate json Error: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "bad incoming json")
	}

	// check if task with "name" already exist
	_, err := h.service.Task.GetTask(task.Name)
	if err == nil {
		log.Warningf("handler CreateTask task already exist: %v", err)
		return echo.NewHTTPError(http.StatusForbidden, "task already exists")
	}

	// create task
	createdTask, err := h.service.Task.CreateTask(&task)
	if err != nil {
		log.Errorf("CreateTask handler Error: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "couldnot create task")
	}
	log.Debugf("handler Task created: %v", createdTask)

	return ctx.JSON(http.StatusCreated, createdTask)
}

// DeleteTask creates new task
func (h *TaskHandler) DeleteTask(ctx echo.Context) error {
	log := h.log
	log.Debug("DeleteTask handler starts")
	defer log.Debug("DeleteTask handler end")

	var task model.DBTask

	name := ctx.QueryParam("name")

	// check if param name exists
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "param name is missing")
	}
	task.Name = name

	// delete task
	err := h.service.Task.DeleteTask(&task)
	if err != nil {
		log.Warningf("handler DeleteTask couldnot delete task: %v", err)

		switch err.Error() {
		case "task doesnot exist":
			return echo.NewHTTPError(http.StatusBadRequest, err)
		case "couldnot delete the task":
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	log.Debugf("handler Task deleted: %s", task.Name)

	return echo.NewHTTPError(http.StatusOK, "task deleted")
}
