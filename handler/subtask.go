package handler

import (
	"job_control_api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateSubTask creates a new subTask
func (h *TaskHandler) CreateSubTask(ctx echo.Context) error {
	log := h.log
	log.Debug("CreateSubTask handler starts")
	defer log.Debug("CreateSubTask handler end")

	var subTask model.DBSubTask

	if err := ctx.Bind(&subTask); err != nil {
		log.Warnf("Bind json Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	// validate subtask
	if err := ctx.Validate(&subTask); err != nil {
		log.Warnf("Validate json Error: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "bad incoming json")
	}
	// TODO: to service!!!!
	// check if subtask with "name" already exist
	_, err := h.service.Task.GetSubTask(subTask.Name)
	if err == nil {
		log.Warnf("CreateSubTask task already exist: %v", err)
		return echo.NewHTTPError(http.StatusForbidden, "subtask already exists")
	}

	// create subtask
	createdSubTask, err := h.service.Task.CreateSubTask(&subTask)
	if err != nil {
		log.Errorf("CreateSubTask handler Error: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "couldnot create subtask")
	}

	log.Debugf("SubTask created: %v", createdSubTask)

	return ctx.JSON(http.StatusCreated, createdSubTask)
}

// DeleteSubTask deletes a subTask
func (h *TaskHandler) DeleteSubTask(ctx echo.Context) error {
	log := h.log
	log.Debug("DeleteSubTask handler starts")
	defer log.Debug("DeleteSubTask handler end")

	var subTask model.DBSubTask

	name := ctx.QueryParam("name")

	// check if param name exists
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "param name is missing")
	}
	subTask.Name = name

	// delete subtask
	err := h.service.Task.DeleteSubTask(&subTask)
	if err != nil {
		log.Warningf("DeleteSubTask handler couldnot delete task: %v", err)

		switch err.Error() {
		case "subtask doesnot exist":
			return echo.NewHTTPError(http.StatusBadRequest, err)
		case "couldnot delete the subtask":
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	log.Debugf("SubTask deleted: %v", subTask.Name)

	return echo.NewHTTPError(http.StatusOK, "subtask deleted")
}

// UpdateSubTask deletes a subTask
func (h *TaskHandler) UpdateSubTask(ctx echo.Context) error {
	log := h.log
	log.Debug("UpdateSubTask handler starts")
	defer log.Debug("UpdateSubTask handler end")

	var subTask model.DBSubTask

	if err := ctx.Bind(&subTask); err != nil {
		log.Warnf("Bind json Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	// validate subtask
	if err := ctx.Validate(&subTask); err != nil {
		log.Warnf("Validate json Error: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "bad incoming json")
	}
	log.Debugf("UpdateSubTask handler update subtask=%v", subTask)

	// update subtask
	newSubtask, err := h.service.Task.UpdateSubTask(&subTask)
	if err != nil {
		log.Warningf("UpdateSubTask handler couldnot delete task: %v", err)
		switch err.Error() {
		case "subtask doesnot exist":
			return echo.NewHTTPError(http.StatusBadRequest, err)
		case "task doesnot exist":
			return echo.NewHTTPError(http.StatusBadRequest, err)
		case "couldnot update subtask":
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	return ctx.JSON(http.StatusOK, newSubtask)

}
