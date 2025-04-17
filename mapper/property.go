package mapper

import (
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

func PropertyModelToResponseDTOMapper(property models.Property) dto.PropertyListDTO {
	return dto.PropertyListDTO{
		ID:           uint32(property.ID),
		Title:        property.Title,
		Purpose:      string(property.Purpose),
		Price:        property.Price,
		PropertyType: property.PropertyType,
		Country: dto.CountryMinimalDTO{
			ID:   uint32(property.Country.ID),
			Name: property.Country.Name,
		},
		Division: dto.DivisionMinimal2DTO{
			ID:   uint32(property.Division.ID),
			Name: property.Division.Name,
		},
		District: dto.DistrictMinimalResponseDTO{
			ID:   uint32(property.District.ID),
			Name: property.District.Name,
		},
		Status:    string(property.Status),
		Address:   property.Address,
		CreatedAt: property.CreatedAt.Format("2006-01-02 15:04:05"),
		Views:     0,
		Inquiries: 0,
	}
}

func PropertyModelToDetailsResponseDTOMapper(property models.Property) dto.PropertyResponseDTO {
	return dto.PropertyResponseDTO{
		ID:           uint32(property.ID),
		Title:        property.Title,
		Purpose:      string(property.Purpose),
		Price:        property.Price,
		PropertyType: property.PropertyType,
		CountryID:    uint32(property.CountryID),
		DivisionID:   uint32(property.DivisionID),
		DistrictID:   uint32(property.DistrictID),
		Status:       string(property.Status),
		Address:      property.Address,
		CreatedAt:    property.CreatedAt.Format("2006-01-02 15:04:05"),
		Description:  property.Description,
		BedRooms:     property.Bedrooms,
		BathRooms:    property.Bathrooms,
		Size:         property.Size,
		BuiltYear:    property.BuiltYear,
	}
}
