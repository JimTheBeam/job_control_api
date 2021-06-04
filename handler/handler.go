package handler

import (
	"job_control_api/config"
	"job_control_api/model"
	"job_control_api/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TaskHandler is a handler for tasks
type TaskHandler struct {
	service *service.Service
	cfg     *config.Config
}

// NewTask ...
func NewTask(service *service.Service, cfg *config.Config) *TaskHandler {
	return &TaskHandler{
		service: service,
		cfg:     cfg,
	}
}

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

	log.Println(task)

	h.service.Task.CreateTask(&task)

	return ctx.JSON(http.StatusCreated, task)
}
