package handler

import (
	"job_control_api/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateTask creates new task
func (h *TaskHandler) CreateTask(ctx echo.Context) error {
	log.Printf("CreateTask handler starts")

	var task model.DBTask

	if err := ctx.Bind(&task); err != nil {
		log.Printf("Bind json Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	// validate task
	if err := ctx.Validate(&task); err != nil {
		log.Printf("Validate json Error: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "bad incoming json")
	}

	// check if task with "name" already exist
	_, err := h.service.Task.GetTask(task.Name)
	if err == nil {
		log.Printf("CreateTask task already exist: %v", err)
		return echo.NewHTTPError(http.StatusForbidden, "task already exists")
	}

	// create task
	createdTask, err := h.service.Task.CreateTask(&task)
	if err != nil {
		log.Printf("CreateTask handler Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	return ctx.JSON(http.StatusCreated, createdTask)
}

// DeleteTask creates new task
func (h *TaskHandler) DeleteTask(ctx echo.Context) error {
	log.Printf("DeleteTask handler starts")

	var task model.DBTask

	if err := ctx.Bind(&task); err != nil {
		log.Printf("Bind json Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	// validate task
	if err := ctx.Validate(&task); err != nil {
		log.Printf("Validate json Error: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "bad incoming json")
	}

	// check if task with "name" exists
	_, err := h.service.Task.GetTask(task.Name)
	if err != nil {
		log.Printf("DeleteTask task doesnot exist: %v", err)
		return echo.NewHTTPError(http.StatusForbidden, "task doesnot exist")
	}

	// delete task
	err = h.service.Task.DeleteTask(&task)
	if err != nil {
		log.Printf("DeleteTask couldnot delete task: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "couldnot delete task")
	}

	return echo.NewHTTPError(http.StatusOK, "task deleted")
}
