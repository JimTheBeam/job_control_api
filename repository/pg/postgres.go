package pg

import (
	"database/sql"
	"fmt"
	"job_control_api/config"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	tasksTable = "tasks"
)

// NewPostgresDB creates connection to postgres database
func NewPostgresDB(cfg *config.Config, log *logrus.Logger) (*sql.DB, error) {
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
		log.Errorf("Database connection: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Errorf("Lost database connection: %v", err)
		return nil, err
	}
	log.Info("Successfully connected to database.")

	return db, nil
}
