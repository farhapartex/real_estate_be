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

func DivisionDtoToModelMapper(request dto.DivisionRequestDTO) models.Division {
	return models.Division{
		Name:      request.Name,
		CountryId: request.CountryId,
	}
}

func DivisionModelToDTOMapper(division models.Division, countryName string, districtCount int64) dto.DivisionResponseDTO {
	return dto.DivisionResponseDTO{
		ID:        division.ID,
		Name:      division.Name,
		Country:   countryName,
		CountryID: division.CountryId,
		Status:    division.Status,
		Districts: districtCount,
	}
}
