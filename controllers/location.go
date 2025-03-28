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

func (c *AuthController) CreateDivision(request dto.DivisionRequestDTO) (*dto.DivisionResponseDTO, error) {
	var country models.Country

	result := c.DB.First(&country, request.CountryId)
	if result.RowsAffected == 0 {
		return nil, errors.New("Country not found")
	}

	modelData := mapper.DivisionDtoToModelMapper(request)
	modelData.Country = country

	tx := c.DB.Begin()
	err := tx.Create(&modelData).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Division creation failed")
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, errors.New("Division creation failed")
	}

	response := mapper.DivisionModelToDTOMapper(modelData, country.Name, 0)

	return &response, nil
}

func (c *AuthController) DivisionList(page, pageSize int) ([]dto.DivisionResponseDTO, int64, error) {
	if c.DB == nil {
		c.DB = config.DB
	}

	var divisions []models.Division
	var total int64

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	result := c.DB.Preload("Country").Order("name ASC").Offset(offset).Limit(pageSize).Find(&divisions)
	if result.Error != nil {
		return nil, 0, errors.New("Failed to fetch divisions")
	}

	c.DB.Model(&models.Division{}).Count(&total)

	var responseDTOs []dto.DivisionResponseDTO
	for _, division := range divisions {
		var districtCount int64
		c.DB.Model(&models.District{}).Where("division_id = ?", division.ID).Count(&districtCount)

		//countryInfo := fmt.Sprintf("{id: %d, name: %s}", division.Country.ID, division.Country.Name)
		dto := mapper.DivisionModelToDTOMapper(division, "", districtCount)
		responseDTOs = append(responseDTOs, dto)
	}

	return responseDTOs, total, nil

}

func (c *AuthController) UpdateDivision(id uint32, request dto.DivisionUpdateRequestDTO) (*dto.DivisionResponseDTO, error) {
	if c.DB == nil {
		c.DB = config.DB
	}
	var division models.Division

	// Only fetch the division once, with preload
	if err := c.DB.Preload("Country").First(&division, id).Error; err != nil {
		return nil, errors.New("Division not found")
	}

	// Check if the country exists if it's being updated
	if request.CountryID != division.CountryId {
		var country models.Country
		if err := c.DB.First(&country, request.CountryID).Error; err != nil {
			return nil, errors.New("New country not found")
		}
	}

	tx := c.DB.Begin()

	// err := tx.Model(&division).Updates(map[string]interface{}{
	// 	"country_id": request.CountryID,
	// 	"name":       request.Name,
	// 	"status":     request.Status,
	// }).Error

	err := tx.Exec("UPDATE divisions SET country_id = ?, name = ?, status = ?, updated_at = NOW() WHERE id = ?",
		request.CountryID, request.Name, request.Status, id).Error

	if err != nil {
		tx.Rollback()
		return nil, errors.New("Failed to update division")
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, errors.New("Failed to update division")
	}

	// After update, fetch the division again with the updated country data
	var updatedDivision models.Division
	if err := c.DB.Preload("Country").First(&updatedDivision, id).Error; err != nil {
		return nil, errors.New("Failed to retrieve updated division")
	}

	var districtCount int64
	c.DB.Model(&models.District{}).Where("division_id = ?", division.ID).Count(&districtCount)

	// Use the updated mapper function
	response := dto.DivisionResponseDTO{
		ID:        updatedDivision.ID,
		Name:      updatedDivision.Name,
		Country:   dto.CountryMinimalDTO{ID: updatedDivision.Country.ID, Name: updatedDivision.Country.Name},
		Status:    updatedDivision.Status,
		Districts: districtCount,
	}

	return &response, nil
}

func (c *AuthController) DeleteDivision(id uint32) error {
	if c.DB == nil {
		c.DB = config.DB
	}
	var division models.Division

	if err := c.DB.First(&division, id).Error; err != nil {
		return errors.New("Division not found")
	}

	var districtCount int64
	c.DB.Model(&models.District{}).Where("division_id = ?", id).Count(&districtCount)
	if districtCount > 0 {
		return errors.New("Cannot delete division with associated districts")
	}

	tx := c.DB.Begin()
	if err := tx.Delete(&division).Error; err != nil {
		tx.Rollback()
		return errors.New("Failed to delete division")
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("Failed to delete division")
	}

	return nil
}
