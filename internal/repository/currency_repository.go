package repository

import (
	"database/sql"
	"fmt"

	"github.com/kalimoldayev02/kmf-task/internal/dto"
)

type CurrencyRepository struct {
	db *sql.DB
}

func newCurrencyRepository(d *sql.DB) *CurrencyRepository {
	return &CurrencyRepository{db: d}
}

func (r *CurrencyRepository) CreateCurrency(currency dto.CreateCurrencDTO) (uint, error) {
	var id uint

	query := fmt.Sprintf("INSERT INTO %s (title, code, value, date) VALUES ($1, UPPER($2), $3, $4) ON CONFLICT DO NOTHING RETURNING id", currencyTable)
	row := r.db.QueryRow(query, currency.Title, currency.Code, currency.Value, currency.Date)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
