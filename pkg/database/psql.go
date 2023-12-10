package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostgresConfig struct {
	Host     string `envconfig:"Host"`
	Port     int    `envconfig:"Port"`
	Username string `envconfig:"Username"`
	Name     string `envconfig:"Name"`
	SSLMode  string `envconfig:"SSLMode"`
	Password string `envconfig:"Password"`
}

func NewPostgresConnection(cfg PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Name, cfg.SSLMode, cfg.Password))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
