package service

import "github.com/kalimoldayev02/kmf-task/app/repository"

type Currency interface {
	Save(date string) bool
}

type Service struct {
	Currency
}

func NewService(r *repository.Respository) *Service {
	return &Service{
		Currency: NewAuthSerive(r.Currency),
	}
}
