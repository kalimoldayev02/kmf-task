package repository

import "database/sql"

const (
	currencyTable = "r_currency"
)

type Currency interface {
}

type Respository struct {
	Currency
}

func NewRepository(db *sql.DB) *Respository {
	return &Respository{
		Currency: NewCurrencyRepository(db),
	}
}
