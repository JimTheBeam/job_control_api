package main

import (
	"flag"
	"fmt"
	"job_control_api/config"
	"job_control_api/handler"
	apiErr "job_control_api/lib/error"
	"job_control_api/lib/validator"
	"job_control_api/logger"
	"job_control_api/repository"
	"job_control_api/repository/pg"
	"job_control_api/service"
	"job_control_api/service/cash"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	err := run()
	if err != nil {
		fmt.Printf("main error: %v", err)
		os.Exit(10)
	}

}

func run() error {
	var (
		cfg      config.Config
		exitCode = 1
	)

	// parse flags
	flg := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	logLvl := flg.String("l", "", "-l <log level>")
	flg.StringVar(logLvl, "log", "", "--log level>")
	flg.Parse(os.Args[1:])

	// ctx := context.Background()

	// init logger
	log := logrus.New()
	logger.InitLogger(log, *logLvl)

	log.Info("app is starting...")

	// TODO: переделать на флаг
	cfgFile := "./config/config.yaml"

	// load config
	if err := config.LoadConfig(cfgFile, &cfg, log); err != nil {
		log.Errorf("Config file unmarshal error: %s", err)
		os.Exit(exitCode)
	}

	// exit code 2
	exitCode++

	// open log file
	if cfg.LogFile != "" {
		log.Infof("Log file is: %s", cfg.LogFile)
		lf, err := os.OpenFile(cfg.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
		if err != nil {
			log.Fatalf("Error opening logfile: %s", err)
			os.Exit(exitCode)
		}
		defer lf.Close()
		log.SetOutput(lf)
	}
	log.Info("log file opened")
	log.Warningln("config:", cfg)

	// exit code 3
	exitCode++

	// Connect to database
	db, err := pg.NewPostgresDB(&cfg, log)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
		os.Exit(exitCode)
	}
	log.Info("connected to database")

	// exit code 4
	exitCode++

	// Init db repository
	repo := repository.NewRepository(db, log)
	log.Info("initialized database repository")

	// init service
	serv := service.NewService(repo, &cfg, log)

	// init handlers
	hand := handler.NewTask(serv, &cfg, log)

	// TODO:
	// init cash
	err = cash.Init(repo, log)
	if err != nil {
		log.Fatal("init cash: %v", err)
		os.Exit(exitCode)
	}
	log.Info("initialized cash")

	// Initialize Echo instance
	e := echo.New()
	e.Validator = validator.NewValidator()

	e.HTTPErrorHandler = apiErr.Error

	// init middleware
	e.Use(middleware.Logger())

	// API v1
	v1 := e.Group("/v1")

	taskRoute := v1.Group("/task")
	// set routes
	// Create a new task.
	// Method - POST
	// Parameter content type application/json
	// request json: {"name": "string", "description": "string"}
	// successful response json: {"name": "string", "description": "string"}
	taskRoute.POST("/", hand.CreateTask)

	// Delete a task.
	// Method - DELETE
	// request with query param: name="task name"
	// successful response json: {"status": 200, "error_name": "OK", "message": "task deleted"}
	taskRoute.DELETE("/delete", hand.DeleteTask)

	subTaskRoute := v1.Group("/subtask")
	// set routes
	// Create a new subtask.
	// Method - POST
	// Parameter content type application/json
	// request json: {"name": "string", "description": "string", "task_name":"string"}TODO:
	// successful response json: {"name": "string", "description": "string"}
	subTaskRoute.POST("/", hand.CreateSubTask)
	// Delete a subtask.
	// Method - DELETE
	// request with query param: name="subtask name"
	// successful response json: {"status": 200, "error_name": "OK", "message": "subtask deleted"}
	subTaskRoute.DELETE("/delete", hand.DeleteSubTask)

	cost := v1.Group("/cost")

	// set routes
	// Create a new cost.
	// Method - POST
	// Parameter content type application/json
	// request json: {"name": "string", "description": "string"} TODO:
	// successful response json: {"name": "string", "cost": "string", "subtask_name": "string"}
	cost.POST("/", hand.CreateCost)

	// TODO: change POST on Delete with query param "name"
	// Delete a cost.
	// Method - DELETE
	// request with query param: name="cost name"
	// successful response json: {"status": 200, "error_name": "OK", "message": "cost deleted"}
	cost.DELETE("/delete", hand.DeleteCost)

	// Start server
	s := &http.Server{
		Addr:         cfg.Server.Addr,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}
