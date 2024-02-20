package repository

import "database/sql"

const (
	currencyTable = "r_currency"
)

type Currency interface {
}

type Repository struct {
	Currency
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Currency: NewCurrencyRepository(db),
	}
}
