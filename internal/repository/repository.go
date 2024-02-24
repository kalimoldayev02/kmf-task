package repository

import (
	"database/sql"

	"github.com/kalimoldayev02/kmf-task/internal/dto"
)

const (
	currencyTable = "currency"
)

type Currency interface {
	CreateCurrency(currency dto.CreateCurrencDTO) (uint, error)
}

type Repository struct {
	Currency
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Currency: newCurrencyRepository(db),
	}
}
