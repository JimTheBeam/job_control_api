package handler

import (
	"job_control_api/config"
	"job_control_api/service"
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
