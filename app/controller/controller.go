package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/kalimoldayev02/kmf-task/app/service"
)

type Controller struct {
	Service   *service.Service
	Validator *validator.Validate
}

func NewController(service *service.Service, validator *validator.Validate) *Controller {
	return &Controller{
		Service:   service,
		Validator: validator,
	}
}