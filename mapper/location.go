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

func DivisionModelToDTOMapper(division models.Division, countryInfo string, districtCount int64) dto.DivisionResponseDTO {
	return dto.DivisionResponseDTO{
		ID:   division.ID,
		Name: division.Name,
		Country: dto.CountryMinimalDTO{
			ID:   division.Country.ID,
			Name: division.Country.Name,
		},
		Status:    division.Status,
		Districts: districtCount,
	}
}

func DistrictDtoToModelMapper(request dto.DistrictRequestDTO, division models.Division) models.District {
	return models.District{
		Name:       request.Name,
		DivisionId: request.DivisionId,
		CountryId:  division.CountryId,
		Status:     true,
	}
}

func DistrictModelToDTOMapper(district models.District) dto.DistrictResponseDTO {
	return dto.DistrictResponseDTO{
		ID:   district.ID,
		Name: district.Name,
		Division: dto.DivisionMinimalDTO{
			ID:   district.Division.ID,
			Name: district.Division.Name,
			Country: dto.CountryMinimalDTO{
				ID:   district.Division.Country.ID,
				Name: district.Division.Country.Name,
			},
		},
		Status: district.Status,
	}
}

func CountryToPublicDTO(country models.Country) dto.PublicCountryDTO {
	return dto.PublicCountryDTO{
		ID:   country.ID,
		Name: country.Name,
		Code: country.Code,
	}
}

func DivisionToPublicDTO(division models.Division) dto.PublicDivisionDTO {
	return dto.PublicDivisionDTO{
		ID:      division.ID,
		Name:    division.Name,
		Country: CountryToPublicDTO(division.Country),
	}
}

func DistrictToPublicDTO(district models.District) dto.PublicDistrictDTO {
	return dto.PublicDistrictDTO{
		ID:   district.ID,
		Name: district.Name,
	}
}

func CreatePaginatedResponse(data interface{}, total int64, page, pageSize int) dto.PaginatedResponse {
	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	return dto.PaginatedResponse{
		Data:       data,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}
