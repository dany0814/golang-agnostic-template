package config

import (
	"context"

	"github.com/go-playground/validator"
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
	MaxDuration  int    `envconfig:"max_age"`
	MaxBackups   int    `envconfig:"max_backups"`
}

var Params Parameters
var Valid *validator.Validate

func LoadConfig(ctx context.Context) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	err = envconfig.Process("TEMP", &Params)
	if err != nil {
		return err
	}

	Valid = validator.New()

	return nil
}
