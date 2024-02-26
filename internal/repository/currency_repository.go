package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/kalimoldayev02/kmf-task/internal/dto"
	"github.com/kalimoldayev02/kmf-task/internal/entity"
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

func (r *CurrencyRepository) GetByDate(date string) ([]entity.Currency, error) {
	parsedDate, err := time.Parse(dateLayout, date)
	if err != nil {
		return nil, err
	}

	var currencies []entity.Currency
	query := fmt.Sprintf("SELECT id, code, value, TO_CHAR(date, 'DD.MM.YYYY') FROM %s WHERE TO_CHAR(date, 'DD-MM-YYYY') = $1 ORDER BY id DESC", currencyTable)

	rows, err := r.db.Query(query, parsedDate.Format(dateLayout))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currency entity.Currency
		var tmpDate string

		err := rows.Scan(&currency.ID, &currency.Code, &currency.Value, &tmpDate)

		if err != nil {
			return nil, err
		}

		currency.Date, err = time.Parse("02.01.2006", tmpDate)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return currencies, nil
}

func (r *CurrencyRepository) GetByDateAndCode(date string, code string) ([]entity.Currency, error) {
	parsedDate, err := time.Parse(dateLayout, date)
	if err != nil {
		return nil, err
	}

	var currencies []entity.Currency
	query := fmt.Sprintf("SELECT id, code, value, TO_CHAR(date, 'DD.MM.YYYY') FROM %s WHERE TO_CHAR(date, 'DD-MM-YYYY') = $1 AND UPPER(code) = $2 ORDER BY id DESC", currencyTable)

	rows, err := r.db.Query(query, parsedDate.Format(dateLayout), strings.ToUpper(code))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currency entity.Currency
		var tmpDate string

		err := rows.Scan(&currency.ID, &currency.Code, &currency.Value, &tmpDate)

		if err != nil {
			return nil, err
		}

		currency.Date, err = time.Parse("02.01.2006", tmpDate)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return currencies, nil
}
