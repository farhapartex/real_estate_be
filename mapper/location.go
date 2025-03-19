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

func CountryModelToDTOMapper(country models.Country, divisionCount int64) dto.CountryResponseDTO {
	return dto.CountryResponseDTO{
		ID:        country.ID,
		Name:      country.Name,
		Code:      country.Code,
		Status:    country.Status,
		Divisions: divisionCount,
	}
}
