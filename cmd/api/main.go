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

	// TODO: change this

	logg := logrus.New()
	logger.InitLogger(logg, *logLvl)

	logg.Info("app is starting...")

	// TODO: переделать на флаг
	cfgFile := "./config/config.yaml"

	// load config
	if err := config.LoadConfig(cfgFile, &cfg); err != nil {
		logg.Errorf("Config file unmarshal error: %s", err)
		os.Exit(exitCode)
	}
	logg.Info("config loaded")

	// exit code 2
	exitCode++

	// open log file
	if cfg.LogFile != "" {
		logg.Infof("Log file is: %s", cfg.LogFile)
		lf, err := os.OpenFile(cfg.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
		if err != nil {
			logg.Fatalf("Error opening logfile: %s", err)
			os.Exit(exitCode)
		}
		defer lf.Close()
		logg.SetOutput(lf)
	}
	logg.Info("log file opened")

	// exit code 3
	exitCode++

	// Connect to database
	db, err := pg.NewPostgresDB(&cfg)
	if err != nil {
		logg.Fatalf("failed to initialize db: %s", err.Error())
		os.Exit(exitCode)
	}
	logg.Info("connected to database")

	// exit code 4
	exitCode++

	// Init db repository
	repo := repository.NewRepository(db)
	logg.Info("initialized database repository")

	// init service
	serv := service.NewService(repo, &cfg)

	// init handlers
	hand := handler.NewTask(serv, &cfg)

	// TODO:
	// init cash
	err = cash.Init(repo)
	if err != nil {
		logg.Fatal("init cash: %v", err)
		os.Exit(exitCode)
	}
	logg.Error("initialized cash")

	// Initialize Echo instance
	e := echo.New()
	e.Validator = validator.NewValidator()

	e.HTTPErrorHandler = apiErr.Error

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
	// Method - POST TODO: change POST on Delete with query param "name"
	// Parameter content type application/json
	// request json: {"name": "string", "description": "string"}
	// successful response json: {"name": "string", "description": "string"} TODO:
	taskRoute.POST("/delete", hand.DeleteTask)

	subTaskRoute := v1.Group("/subtask")
	// set routes
	// Create a new subtask.
	// Method - POST
	// Parameter content type application/json
	// request json: {"name": "string", "description": "string"}TODO:
	// successful response json: {"name": "string", "description": "string"}
	subTaskRoute.POST("/", hand.CreateSubTask)
	// Delete a task.
	// Method - POST TODO: change POST on Delete with query param "name"
	// Parameter content type application/json
	// request json: {"name": "string", "description": "string"}TODO:
	// successful response json: {"name": "string", "description": "string"} TODO:
	subTaskRoute.POST("/delete", hand.DeleteSubTask)

	cost := v1.Group("/cost")

	cost.POST("/", hand.CreateCost)

	// TODO: change POST on Delete with query param "name"
	cost.POST("/delete", hand.DeleteCost)

	// Start server
	s := &http.Server{
		Addr:         cfg.Server.Addr,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}
