package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kalimoldayev02/kmf-task/internal/db"
	handler "github.com/kalimoldayev02/kmf-task/internal/delivery/http"
	"github.com/kalimoldayev02/kmf-task/internal/repository"
	"github.com/kalimoldayev02/kmf-task/internal/service"
	"github.com/kalimoldayev02/kmf-task/pkg/config"
	"github.com/kalimoldayev02/kmf-task/pkg/utils"
)

func Run(cfg *config.Config) {
	// DB
	db, err := db.NewDb(cfg)
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
	handler.NotFoundRoute()

	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", cfg.HttpServer.Port), handler.GetRouter()))
}
