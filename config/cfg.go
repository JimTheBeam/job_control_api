package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Config is a config
type Config struct {
	Server  ServerConfig `yaml:"server_config"`
	DB      DBConfig     `yaml:"database_config"`
	LogFile string       `yaml:"log_path"`
}

// ServerConfig is a config for a server
type ServerConfig struct {
	Addr         string        `yaml:"http_addr"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

// DBConfig is a config for a postgres database
type DBConfig struct {
	Username string `yaml:"username"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
	SSLMode  string `yaml:"ssl_mode"`
	Password string `yaml:"db_password"`
}

// LoadCfg - open config file and put config to cfg.Config struct
func LoadConfig(path string, cfg *Config, log *logrus.Logger) error {
	log.Debug("Loading config")
	defer log.Debug("Config loaded")

	log.Debug("Config file: %s", path)
	cfgData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Errorf("Unable to read config file: %v", err)
		return err
	}

	err = yaml.Unmarshal(cfgData, &cfg)
	if err != nil {
		log.Errorf("Unable to unmarshal config data: %v", err)
		return err
	}

	// print config
	configBytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Warn(err)
	}
	fmt.Println("Configuration:", string(configBytes))
	log.Debug("Configuration:", string(configBytes))

	return nil
}
