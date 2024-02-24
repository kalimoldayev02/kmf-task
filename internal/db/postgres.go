package db

import (
	"database/sql"
	"fmt"

	"github.com/kalimoldayev02/kmf-task/pkg/config"
)

func newPOstgresDb(cfg *config.Database) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", cfg.Host, cfg.Username, cfg.DBName, cfg.Password)

	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
