package repository

import (
	"database/sql"

	"github.com/kalimoldayev02/kmf-task/internal/dto"
	"github.com/kalimoldayev02/kmf-task/internal/entity"
)

const (
	currencyTable = "currency"
	dateLayout    = "02-01-2006"
)

type Currency interface {
	CreateCurrency(currency dto.CreateCurrencDTO) (uint, error)
	GetByDate(date string) ([]entity.Currency, error)
	GetByDateAndCode(date string, code string) ([]entity.Currency, error)
}

type Repository struct {
	Currency
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Currency: newCurrencyRepository(db),
	}
}
