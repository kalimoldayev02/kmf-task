package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	handler "github.com/kalimoldayev02/kmf-task/internal/delivery/http"
	"github.com/kalimoldayev02/kmf-task/internal/repository"
	"github.com/kalimoldayev02/kmf-task/internal/service"
	"github.com/kalimoldayev02/kmf-task/pkg/config"
	"github.com/kalimoldayev02/kmf-task/pkg/utils"
)

func Run(cfg *config.Config) {
	// DB
	db, err := initDB(cfg)
	if err != nil {
		log.Printf("error on init db: %s", err.Error())
	}

	// repository
	repository := repository.NewRepository(db)

	// service
	service := service.NewService(repository)

	// validator
	validator := utils.NewValidator()

	// handler
	handler := handler.NewHandler(service, validator)
	handler.PublicHandler()

	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", cfg.HttpServer.Port), handler.GetRouter()))
}

func initDB(c *config.Config) (*sql.DB, error) {
	var dsn string
	cfg := c.Database

	switch cfg.Driver {
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", cfg.Host, cfg.Username, cfg.DBName, cfg.Password)
	default:
		return nil, fmt.Errorf("undefined database driver: %s", cfg.Driver)
	}

	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	defer db.Close()

	return db, nil
}
