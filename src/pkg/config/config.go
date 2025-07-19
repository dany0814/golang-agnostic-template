package config

import (
	"context"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Parameters struct {
	// Overview
	Environment string `envconfig:"environment"`
	// Logger
	Level        string `envconfig:"level"`
	TraceLevel   string `envconfig:"trace_level"`
	FileEnabled  bool   `envconfig:"file_enabled"`
	FileSize     int    `envconfig:"file_size"`
	FilePath     string `envconfig:"file_path"`
	FileCompress bool   `envconfig:"file_compress"`
	MaxDuration  int    `envconfig:"max_duration"`
	MaxBackups   int    `envconfig:"max_backups"`
	// Web Server
	WebHost         string        `envconfig:"web_host" default:"0.0.0.0"`
	WebPort         int           `envconfig:"web_port"`
	ShutdownTimeout time.Duration `envconfig:"web_shutdowntimeout" default:"20s"`
	// Database
	DBHost      string `envconfig:"db_host"`
	DBPort      int    `envconfig:"db_port"`
	DBNamespace string `envconfig:"db_namespace"`
	DBDatabase  string `envconfig:"db_database"`
	DBUser      string `envconfig:"db_user"`
	DBPassword  string `envconfig:"db_password"`
	// Token
	TokenSecret string `envconfig:"secret_key"`
	ExpiredHour int64  `envconfig:"expired_hour"`
}

var Params Parameters

func LoadConfig(ctx context.Context) error {
	err := godotenv.Load(".env.app")
	if err != nil {
		return err
	}

	err = envconfig.Process("TEMP", &Params)
	if err != nil {
		return err
	}
	return nil
}
