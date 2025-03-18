package mapper

import (
	"github.com/farhapartex/real_estate_be/dto"
	"github.com/farhapartex/real_estate_be/models"
)

func CountryDtoToModelMapper(request dto.CountryRequestDTO) models.Country {
	return models.Country{
		Name: request.Name,
		Code: request.Code,
	}
}
