package domain

import (
	"github.com/dmytrodemianchuk/bank-transactions/pkg/database"
)

type HTTPConfig struct {
	Host string
	Port string
}

type Config struct {
	Postgres database.PostgresConfig
	HTTP     HTTPConfig
}
