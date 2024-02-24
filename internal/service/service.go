package service

import (
	"github.com/kalimoldayev02/kmf-task/internal/dto"
	"github.com/kalimoldayev02/kmf-task/internal/repository"
)

type Currency interface {
	Create(date string) bool
	GetByDate(date string) ([]dto.ResponseCurrencyDTO, error)
	GetByDateAndCode(date string, code string) ([]dto.ResponseCurrencyDTO, error)
}

type Service struct {
	Currency
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Currency: newCurrencyService(r.Currency),
	}
}
