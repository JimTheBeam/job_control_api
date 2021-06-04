package main

import (
	"flag"
	"job_control_api/config"
	"job_control_api/handler"
	apiErr "job_control_api/lib/error"
	"job_control_api/lib/validator"
	"job_control_api/repository"
	"job_control_api/repository/pg"
	"job_control_api/service"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

}

func run() error {
	var (
		cfg      config.Config
		exitCode = 1
	)

	// parse flags
	flg := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	logFile := flg.String("l", "", "-l <path to log file>")
	flg.StringVar(logFile, "log", "", "--log <path to log file>")
	flg.Parse(os.Args[1:])

	// ctx := context.Background()

	// open log file
	if *logFile != "" {
		log.Printf("Log file is: %s", *logFile)
		lf, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
		if err != nil {
			log.Printf("Error opening logfile: %s", err)
			_, usage := flag.UnquoteUsage(flg.Lookup("o"))
			log.Printf("Usage: %v", usage)
			os.Exit(exitCode)
		}
		defer lf.Close()
		log.SetOutput(lf)
	}

	log.Println("app is starting...")

	// TODO: переделать на флаг
	cfgFile := "./config/config.yaml"

	// load config
	if err := config.LoadConfig(cfgFile, &cfg); err != nil {
		log.Printf("Config file unmarshal error: %s", err)
	}

	// Connect to database
	db, err := pg.NewPostgresDB(&cfg)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	// Init db repository
	repo := repository.NewRepository(db)
	log.Printf("initialized database repository")

	// init service
	serv := service.NewService(repo, &cfg)

	// init handlers
	hand := handler.NewTask(serv, &cfg)

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

	// Create a new task.
	// Method - POST
	// Parameter content type application/json
	// request json: {"name": "string", "description": "string"}
	// successful response json: {"name": "string", "description": "string"} TODO:
	taskRoute.POST("/delete", hand.DeleteTask)

	// Start server
	s := &http.Server{
		Addr:         cfg.Server.Addr,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}
