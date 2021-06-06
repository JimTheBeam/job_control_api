package handler

import (
	"job_control_api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateCost creates a new cost for a subtask
func (h *TaskHandler) CreateCost(ctx echo.Context) error {
	log := h.log
	log.Debug("CreateCost handler starts")
	defer log.Debug("CreateCost handler starts")

	// TODO:
	var cost model.DBCost

	if err := ctx.Bind(&cost); err != nil {
		log.Errorf("Bind json Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	// validate cost
	if err := ctx.Validate(&cost); err != nil {
		log.Errorf("Validate json Error: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "bad incoming json")
	}

	// check if cost with "subtask name" already exist
	_, err := h.service.Task.GetCost(cost.SubTaskName)
	if err == nil {
		log.Errorf("CreateCost Cost already exist: %v", err)
		return echo.NewHTTPError(http.StatusForbidden, "cost already exists")
	}

	// create subtask
	createdCost, err := h.service.Task.CreateCost(&cost)
	if err != nil {
		log.Errorf("CreateCost handler Error: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "couldnot create cost")
	}

	log.Debugf("Cost created: %v", createdCost)

	return ctx.JSON(http.StatusCreated, createdCost)
}

// DeleteSubTask deletes a subTask
func (h *TaskHandler) DeleteCost(ctx echo.Context) error {
	log := h.log
	log.Debug("DeleteCost handler starts")
	defer log.Debug("DeleteCost handler end")

	// TODO:
	var cost model.DBCost

	if err := ctx.Bind(&cost); err != nil {
		log.Errorf("Bind json Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	// validate cost
	if err := ctx.Validate(&cost); err != nil {
		log.Errorf("Validate json Error: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "bad incoming json")
	}

	// check if cost with "subtask name" exists
	_, err := h.service.Task.GetSubTask(cost.SubTaskName)
	if err != nil {
		log.Errorf("DeleteCost subtask doesnot exist: %v", err)
		return echo.NewHTTPError(http.StatusForbidden, "subtask doesnot exist")
	}

	// delete cost
	err = h.service.Task.DeleteCost(&cost)
	if err != nil {
		log.Errorf("DeleteCost couldnot delete cost: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "couldnot delete cost")
	}

	log.Debugf("Cost deleted: %v", cost)

	return echo.NewHTTPError(http.StatusOK, "cost deleted")
}
