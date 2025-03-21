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

	response := mapper.CountryModelToDTOMapper(newCountry, 0)

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

	offset := (page - 1) * pageSize

	result := c.DB.Order("name ASC").Offset(offset).Limit(pageSize).Find(&countries)
	if result.Error != nil {
		return nil, 0, errors.New("Failed to fetch countries")
	}

	// Map models to DTOs
	var responseDTOs []dto.CountryResponseDTO
	for _, country := range countries {
		var divisionCount int64
		c.DB.Model(&models.Division{}).Where("country_id = ?", country.ID).Count(&divisionCount)

		dto := mapper.CountryModelToDTOMapper(country, divisionCount)
		responseDTOs = append(responseDTOs, dto)
	}

	return responseDTOs, total, nil
}

func (c *AuthController) UpdateCountry(id uint32, request dto.CountryUpdateRequestDTO) (*dto.CountryResponseDTO, error) {
	if c.DB == nil {
		c.DB = config.DB
	}
	var country models.Country

	if err := c.DB.First(&country, id).Error; err != nil {
		return nil, errors.New("Cuontry not found")
	}
	if country.Code != request.Code {
		var existingCountry models.Country
		result := c.DB.Where("code = ? AND id != ?", request.Code, id).First(&existingCountry)
		if result.RowsAffected > 0 {
			return nil, errors.New("Country with this code already exists")
		}
	}

	tx := c.DB.Begin()
	country.Name = request.Name
	country.Code = request.Code
	country.Status = request.Status

	err := tx.Save(&country).Error
	if err != nil {
		return nil, errors.New("Failed to update country")
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, errors.New("Failed to update country")
	}

	var divisionCount int64
	c.DB.Model(&models.Division{}).Where("country_id = ?", country.ID).Count(&divisionCount)

	response := mapper.CountryModelToDTOMapper(country, divisionCount)

	return &response, nil
}

func (c *AuthController) DeleteCountry(id uint32) error {
	if c.DB == nil {
		c.DB = config.DB
	}
	var country models.Country

	if err := c.DB.First(&country, id).Error; err != nil {
		return errors.New("Cuontry not found")
	}

	var divisionCount int64
	c.DB.Model(&models.Division{}).Where("country_id = ?", id).Count(&divisionCount)
	if divisionCount > 0 {
		return errors.New("Cannot delete country with associated divisions")
	}

	tx := c.DB.Begin()
	if err := tx.Delete(&country).Error; err != nil {
		tx.Rollback()
		return errors.New("Failed to delete country")
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("Failed to delete country")
	}

	return nil
}

func (c *AuthController) CreateDivision(request dto.CountryRequestDTO) (*dto.CountryResponseDTO, error) {
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

	response := mapper.CountryModelToDTOMapper(newCountry, 0)

	return &response, nil
}
