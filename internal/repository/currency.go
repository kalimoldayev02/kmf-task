package repository

import "database/sql"

type CurrencyRepository struct {
	db *sql.DB
}

func NewCurrencyRepository(d *sql.DB) *CurrencyRepository {
	return &CurrencyRepository{db: d}
}
