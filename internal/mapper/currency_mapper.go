package mapper

import (
	"time"

	"github.com/kalimoldayev02/kmf-task/internal/dto"
	"github.com/kalimoldayev02/kmf-task/internal/entity"
)

func MapRequestToCreate(request dto.RequestCurrencyDTO, date time.Time) dto.CreateCurrencDTO {
	return dto.CreateCurrencDTO{
		Title: request.FullName,
		Code:  request.Title,
		Value: request.Description,
		Date:  date,
	}
}

func MapEntityToResponse(entity entity.Currency) dto.ResponseCurrencyDTO {
	return dto.ResponseCurrencyDTO{
		ID:    entity.ID,
		Code:  entity.Code,
		Value: entity.Value,
		Date:  entity.Date.Format("15.09.2013"),
	}
}
