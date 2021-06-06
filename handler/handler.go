package handler

import (
	"job_control_api/config"
	"job_control_api/service"

	"github.com/sirupsen/logrus"
)

// TaskHandler is a handler for tasks
type TaskHandler struct {
	service *service.Service
	cfg     *config.Config
	log     *logrus.Logger
}

// NewTask ...
func NewTask(service *service.Service, cfg *config.Config, log *logrus.Logger) *TaskHandler {
	return &TaskHandler{
		service: service,
		cfg:     cfg,
		log:     log,
	}
}
