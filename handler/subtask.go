package handler

import (
	"job_control_api/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateSubTask creates a new subTask
func (h *TaskHandler) CreateSubTask(ctx echo.Context) error {
	var subTask model.DBSubTask

	if err := ctx.Bind(&subTask); err != nil {
		log.Printf("Bind json Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	// validate subtask
	if err := ctx.Validate(&subTask); err != nil {
		log.Printf("Validate json Error: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "bad incoming json")
	}

	// check if subtask with "name" already exist
	_, err := h.service.Task.GetSubTask(subTask.Name)
	if err == nil {
		log.Printf("CreateSubTask task already exist: %v", err)
		return echo.NewHTTPError(http.StatusForbidden, "subtask already exists")
	}

	// create subtask
	createdSubTask, err := h.service.Task.CreateSubTask(&subTask)
	if err != nil {
		log.Printf("CreateSubTask handler Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	return ctx.JSON(http.StatusCreated, createdSubTask)
}

// DeleteSubTask deletes a subTask
func (h *TaskHandler) DeleteSubTask(ctx echo.Context) error {
	log.Printf("DeleteSubTask handler starts")

	var subTask model.DBSubTask

	if err := ctx.Bind(&subTask); err != nil {
		log.Printf("Bind json Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	// validate subtask
	if err := ctx.Validate(&subTask); err != nil {
		log.Printf("Validate json Error: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "bad incoming json")
	}

	// check if subtask with "name" exists
	_, err := h.service.Task.GetSubTask(subTask.Name)
	if err != nil {
		log.Printf("DeleteSubTask subtask doesnot exist: %v", err)
		return echo.NewHTTPError(http.StatusForbidden, "subtask doesnot exist")
	}

	// delete subtask
	err = h.service.Task.DeleteSubTask(&subTask)
	if err != nil {
		log.Printf("DeleteSubTask couldnot delete task: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "couldnot delete subtask")
	}

	return echo.NewHTTPError(http.StatusOK, "subtask deleted")
}
