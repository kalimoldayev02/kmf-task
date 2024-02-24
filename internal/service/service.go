package service

import (
	"github.com/kalimoldayev02/kmf-task/internal/repository"
)

type Currency interface {
	Create(date string) bool
}

type Service struct {
	Currency
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Currency: newCurrencyService(r.Currency),
	}
}
