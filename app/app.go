package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/kalimoldayev02/kmf-task/app/controller"
	"github.com/kalimoldayev02/kmf-task/app/repository"
	"github.com/kalimoldayev02/kmf-task/app/service"
	"github.com/kalimoldayev02/kmf-task/pkg/config"
	"github.com/kalimoldayev02/kmf-task/pkg/route"
	"github.com/kalimoldayev02/kmf-task/pkg/utils"
)

func Run(c *config.Config) {
	// DB
	db, err := initDB(c)
	if err != nil {
		log.Printf("error on init db: %s", err.Error())
	}

	// Repository
	repository := repository.NewRepository(db)

	// Service
	service := service.NewService(repository)

	// Validator
	validator := utils.NewValidator()

	// Controller
	controller := controller.NewController(service, validator)

	// Router
	router := route.NewRouter()
	route.PublicRoutes(router, controller)
	route.NotFoundRoute(router)

	http.ListenAndServe(fmt.Sprintf(":%d", c.HttpServer.Port), router)
}

func initDB(c *config.Config) (*sql.DB, error) {
	var dsn string
	cfg := c.Database

	switch cfg.Driver {
	case "mssql":
		dsn = fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d", cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port)

	}

	return sql.Open(cfg.Driver, dsn)
}
