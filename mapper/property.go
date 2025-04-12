package mapper

import (
	"time"

	"github.com/farhapartex/real_estate_be/dto"
	"github.com/farhapartex/real_estate_be/models"
)

func PropertyDtoToModelMapper(request dto.PropertyRequestDTO, userID uint) models.Property {
	return models.Property{
		OwnerID:      userID,
		Title:        request.Title,
		Purpose:      models.PropertyString(request.Purpose),
		Price:        request.Price,
		Status:       models.StatusDraft, // Default status is draft
		PropertyType: request.PropertyType,
		Bedrooms:     request.Bedrooms,
		Bathrooms:    request.Bathrooms,
		Size:         request.Size,
		BuiltYear:    request.BuiltYear,
		CountryID:    request.CountryID,
		DivisionID:   request.DivisionID,
		DistrictID:   request.DistrictID,
		Address:      request.Address,
		Description:  request.Description,
	}
}

func PropertyModelToResponseDTOMapper(property models.Property) dto.PropertyResponseDTO {
	return dto.PropertyResponseDTO{
		ID:           uint32(property.ID),
		Title:        property.Title,
		Purpose:      string(property.Purpose),
		Price:        property.Price,
		Status:       string(property.Status),
		PropertyType: property.PropertyType,
		BedRooms:     property.Bedrooms,
		BathRooms:    property.Bathrooms,
		Size:         property.Size,
		BuiltYear:    property.BuiltYear,
		Country: dto.CountryResponseDTO{
			ID:   uint32(property.Country.ID),
			Name: property.Country.Name,
			Code: property.Country.Code,
		},
		Division: dto.DivisionResponseDTO{
			ID:   uint32(property.Division.ID),
			Name: property.Division.Name,
		},
		District: dto.DistrictResponseDTO{
			ID:   uint32(property.District.ID),
			Name: property.District.Name,
		},
		Address:     property.Address,
		Description: property.Description,
		CreatedAt:   property.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   property.UpdatedAt.Format(time.RFC3339),
	}
}

func PropertyListToResponseDTOMapper(properties []models.Property, total int64, page, perPage int) dto.PropertyListResponseDTO {
	propertyDTOs := make([]dto.PropertyResponseDTO, len(properties))

	for i, property := range properties {
		propertyDTOs[i] = PropertyModelToResponseDTOMapper(property)
	}

	return dto.PropertyListResponseDTO{
		Properties: propertyDTOs,
		Total:      total,
		Page:       page,
		PerPage:    perPage,
	}
}
