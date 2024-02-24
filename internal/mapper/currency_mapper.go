package mapper

import (
	"time"

	"github.com/kalimoldayev02/kmf-task/internal/dto"
)

func MapRequestToCreate(request dto.RequestCurrencyDTO, date time.Time) dto.CreateCurrencDTO {
	return dto.CreateCurrencDTO{
		Title: request.FullName,
		Code:  request.Title,
		Value: request.Description,
		Date:  date,
	}
}
