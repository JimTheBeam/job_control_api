package pg

import (
	"database/sql"
	"fmt"
	"job_control_api/config"
	"log"

	_ "github.com/lib/pq"
)

const (
	tasksTable = "tasks"
)

// NewPostgresDB creates connection to postgres database
func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	// connect to database
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Username,
		cfg.DB.DBName,
		cfg.DB.Password,
		cfg.DB.SSLMode,
	),
	)
	if err != nil {
		log.Printf("Database connection: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Lost database connection: %v", err)
		return nil, err
	}
	log.Println("Successfully connected to database.")

	return db, nil
}
