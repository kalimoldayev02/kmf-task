package db

import (
	"database/sql"
	"fmt"

	"github.com/kalimoldayev02/kmf-task/pkg/config"
)

func NewDb(cfg *config.Config) (*sql.DB, error) {
	switch cfg.Database.Driver {
	case "postgres":
		return newPOstgresDb(&cfg.Database)
	default:
		return nil, fmt.Errorf("undefined database driver: %s", cfg.Database.Driver)
	}
}
