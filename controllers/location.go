package controllers

import (
	"errors"

	"github.com/farhapartex/real_estate_be/config"
	"github.com/farhapartex/real_estate_be/dto"
	"github.com/farhapartex/real_estate_be/mapper"
	"github.com/farhapartex/real_estate_be/models"
)

func (c *AuthController) CreateCountry(request dto.CountryRequestDTO) (*dto.CountryResponseDTO, error) {
	var country models.Country

	result := c.DB.Where("code = ?", request.Code).First(&country)
	if result.RowsAffected > 0 {
		return nil, errors.New("Cuontry exists")
	}

	newCountry := mapper.CountryDtoToModelMapper(request)
	tx := c.DB.Begin()
	err := tx.Create(&newCountry).Error
	if err != nil {
		return nil, errors.New("Country creation failed")
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, errors.New("Country creation failed")
	}

	response := mapper.CountryModelToDTOMapper(newCountry)

	return &response, nil
}

func (c *AuthController) ListCountries(page, pageSize int) ([]dto.CountryResponseDTO, int64, error) {
	// Ensure DB is initialized
	if c.DB == nil {
		c.DB = config.DB
	}

	var countries []models.Country
	var total int64

	// Count total records
	c.DB.Model(&models.Country{}).Count(&total)

	// Set default pagination if not provided
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Get paginated records
	result := c.DB.Order("name ASC").Offset(offset).Limit(pageSize).Find(&countries)
	if result.Error != nil {
		return nil, 0, errors.New("Failed to fetch countries")
	}

	// Map models to DTOs
	var responseDTOs []dto.CountryResponseDTO
	for _, country := range countries {
		dto := mapper.CountryModelToDTOMapper(country)
		responseDTOs = append(responseDTOs, dto)
	}

	return responseDTOs, total, nil
}
