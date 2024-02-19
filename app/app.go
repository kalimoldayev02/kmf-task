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

	// repository
	repository := repository.NewRepository(db)

	// service
	service := service.NewService(repository)

	// validator
	validator := utils.NewValidator()

	// controller
	controller := controller.NewController(service, validator)

	// router
	router := route.NewRouter()
	route.PublicRoutes(router, controller)
	route.NotFoundRoute(router)

	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", c.HttpServer.Port), router))
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
