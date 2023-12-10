package config

import (
	"errors"
	"runtime"

	"github.com/dmytrodemianchuk/bank-transactions/internal/domain"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func Init() (*domain.Config, error) {
	var cfg domain.Config

	if err := setFromEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setFromEnv(cfg *domain.Config) error {

	if runtime.GOOS == "windows" {
		// windows specific code here - for windows debug env
		err := godotenv.Load()
		if err != nil {
			return errors.New("error loading .env file")
		}
	}

	if err := envconfig.Process("db", &cfg.Postgres); err != nil {
		return err
	}
	if err := envconfig.Process("http", &cfg.HTTP); err != nil {
		return err
	}

	return nil
}
