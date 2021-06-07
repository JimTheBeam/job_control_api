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

	// create cost
	createdCost, err := h.service.Task.CreateCost(&cost)
	if err != nil {
		log.Errorf("CreateCost handler Error: %v", err)

		switch err.Error() {
		case "cost already exists":
			return echo.NewHTTPError(http.StatusBadRequest, err)
		case "subtask doesnot exist":
			return echo.NewHTTPError(http.StatusBadRequest, err)
		case "couldnot create cost":
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	log.Debugf("Cost created: %v", createdCost)

	return ctx.JSON(http.StatusCreated, createdCost)
}

// DeleteSubTask deletes a cost with name
func (h *TaskHandler) DeleteCost(ctx echo.Context) error {
	log := h.log
	log.Debug("DeleteCost handler starts")
	defer log.Debug("DeleteCost handler end")

	var cost model.DBCost

	name := ctx.QueryParam("name")
	log.Debugf("DeleteCost handler: deleting=%s", name)
	// check if param name exists
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "param name is missing")
	}

	cost.Name = name

	// delete cost
	err := h.service.Task.DeleteCost(&cost)
	if err != nil {
		log.Warningf("DeleteSubTask handler couldnot delete task: %v", err)

		switch err.Error() {
		case "cost doesnot exist":
			return echo.NewHTTPError(http.StatusBadRequest, err)
		case "couldnot delete cost":
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	log.Debugf("Cost deleted: %v", cost)

	return echo.NewHTTPError(http.StatusOK, "cost deleted")
}

// UpdateCost update a cost with name
func (h *TaskHandler) UpdateCost(ctx echo.Context) error {
	log := h.log
	log.Debug("UpdateCost handler starts")
	defer log.Debug("UpdateCost handler end")

	var cost model.DBCost

	if err := ctx.Bind(&cost); err != nil {
		log.Warningf("Bind json Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad incoming json")
	}

	// validate subtask
	if err := ctx.Validate(&cost); err != nil {
		log.Warningf("Validate json Error: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "bad incoming json")
	}

	log.Debugf("UpdateCost handler updatedcost='%v'", cost)

	// update cost
	updatedCost, err := h.service.Task.UpdateCost(&cost)
	if err != nil {
		log.Warningf("UpdateCost handler couldnot update cost", err)

		switch err.Error() {
		case "cost doesnot exist":
			return echo.NewHTTPError(http.StatusBadRequest, err)

		case "subtask doesnot exist":
			return echo.NewHTTPError(http.StatusBadRequest, err)

		case "couldnot update cost":
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}
	log.Debugf("Cost updated: %v", updatedCost)

	return ctx.JSON(http.StatusOK, updatedCost)
}
