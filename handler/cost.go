package handler

import (
	"job_control_api/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateCost creates a new cost for a subtask
func (h *TaskHandler) CreateCost(ctx echo.Context) error {
	log.Printf("CreateCost handler starts")

	// TODO:
	var cost model.DBCost

	if err := ctx.Bind(&cost); err != nil {
		log.Printf("Bind json Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	// validate cost
	if err := ctx.Validate(&cost); err != nil {
		log.Printf("Validate json Error: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "bad incoming json")
	}

	// check if cost with "subtask name" already exist
	_, err := h.service.Task.GetCost(cost.SubTaskName)
	if err == nil {
		log.Printf("CreateCost Cost already exist: %v", err)
		return echo.NewHTTPError(http.StatusForbidden, "cost already exists")
	}

	// create subtask
	createdCost, err := h.service.Task.CreateCost(&cost)
	if err != nil {
		log.Printf("CreateCost handler Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	return ctx.JSON(http.StatusCreated, createdCost)
}

// DeleteSubTask deletes a subTask
func (h *TaskHandler) DeleteCost(ctx echo.Context) error {
	log.Printf("DeleteCost handler starts")

	// TODO:
	var cost model.DBCost

	if err := ctx.Bind(&cost); err != nil {
		log.Printf("Bind json Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	// validate cost
	if err := ctx.Validate(&cost); err != nil {
		log.Printf("Validate json Error: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "bad incoming json")
	}

	// check if cost with "subtask name" exists
	_, err := h.service.Task.GetSubTask(cost.SubTaskName)
	if err != nil {
		log.Printf("DeleteCost subtask doesnot exist: %v", err)
		return echo.NewHTTPError(http.StatusForbidden, "subtask doesnot exist")
	}

	// delete cost
	err = h.service.Task.DeleteCost(&cost)
	if err != nil {
		log.Printf("DeleteCost couldnot delete cost: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "couldnot delete cost")
	}

	return echo.NewHTTPError(http.StatusOK, "cost deleted")
}
