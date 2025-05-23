package controllers

import (
	"errors"
	"strconv"

	"github.com/farhapartex/real_estate_be/config"
	"github.com/farhapartex/real_estate_be/dto"
	"github.com/farhapartex/real_estate_be/mapper"
	"github.com/farhapartex/real_estate_be/models"
	"github.com/gin-gonic/gin"
)

func GetPaginationParams(c *gin.Context) (int, int) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	return page, pageSize
}

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

func (c *AuthController) CreateDistrict(request dto.DistrictRequestDTO) (*dto.DistrictResponseDTO, error) {
	var division models.Division

	result := c.DB.Preload("Country").First(&division, request.DivisionId)
	if result.RowsAffected == 0 {
		return nil, errors.New("Division not exists")
	}

	district := mapper.DistrictDtoToModelMapper(request, division)
	tx := c.DB.Begin()
	err := tx.Create(&district).Error
	if err != nil {
		return nil, errors.New("District creation failed")
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, errors.New("District commit creation failed")
	}

	district.Division = division
	district.Country = division.Country

	response := mapper.DistrictModelToDTOMapper(district)

	return &response, nil
}

func (c *AuthController) DistrictList(page, pageSize int) ([]dto.DistrictResponseDTO, int64, error) {
	if c.DB == nil {
		c.DB = config.DB
	}

	var districts []models.District
	var total int64

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	result := c.DB.Preload("Country").Preload("Division").Preload("Division.Country").Order("name ASC").Offset(offset).Limit(pageSize).Find(&districts)
	if result.Error != nil {
		return nil, 0, errors.New("Failed to fetch districts")
	}

	c.DB.Model(&models.District{}).Count(&total)

	var responseDTOs []dto.DistrictResponseDTO
	for _, district := range districts {
		dto := mapper.DistrictModelToDTOMapper(district)
		responseDTOs = append(responseDTOs, dto)
	}

	return responseDTOs, total, nil

}

func (c *AuthController) UpdateDistrict(id uint32, request dto.DistrictUpdateRequestDTO) (*dto.DistrictResponseDTO, error) {
	if c.DB == nil {
		c.DB = config.DB
	}
	var district models.District

	// Only fetch the division once, with preload
	if err := c.DB.Preload("Country").Preload("Division").Preload("Division.Country").First(&district, id).Error; err != nil {
		return nil, errors.New("District not found")
	}

	// Check if the country exists if it's being updated
	var division models.Division
	if request.DivisionId != district.DivisionId {
		if err := c.DB.Preload("Country").First(&division, request.DivisionId).Error; err != nil {
			return nil, errors.New("New division not found")
		}
	}

	tx := c.DB.Begin()

	err := tx.Exec("UPDATE districts SET division_id = ?, name = ?, status = ?, updated_at = NOW() WHERE id = ?",
		request.DivisionId, request.Name, request.Status, id).Error

	if err != nil {
		tx.Rollback()
		return nil, errors.New("Failed to update district")
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, errors.New("Failed to update district")
	}

	// After update, fetch the division again with the updated country data
	var updatedDistrict models.District
	if err := c.DB.Preload("Division").Preload("Division.Country").First(&updatedDistrict, id).Error; err != nil {
		return nil, errors.New("Failed to retrieve updated division")
	}

	c.DB.Preload("Country").First(&division, request.DivisionId)

	updatedDistrict.Division = division
	updatedDistrict.Country = division.Country

	response := mapper.DistrictModelToDTOMapper(updatedDistrict)

	return &response, nil
}

func (c *AuthController) DeleteDistrict(id uint32) error {
	if c.DB == nil {
		c.DB = config.DB
	}
	var district models.District

	if err := c.DB.First(&district, id).Error; err != nil {
		return errors.New("District not found")
	}

	tx := c.DB.Begin()
	if err := tx.Delete(&district).Error; err != nil {
		tx.Rollback()
		return errors.New("Failed to delete District")
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("Failed to delete District")
	}

	return nil
}

func (pc *AuthController) GetCountries(page, pageSize int) (*dto.PaginatedResponse, error) {
	// Ensure DB is initialized
	if pc.DB == nil {
		pc.DB = config.DB
	}

	offset := (page - 1) * pageSize

	var countries []models.Country
	var total int64

	// Count total active records
	pc.DB.Model(&models.Country{}).Where("status = ?", true).Count(&total)

	// Get records with pagination
	result := pc.DB.Where("status = ?", true).
		Order("name ASC").
		Offset(offset).
		Limit(pageSize).
		Find(&countries)

	if result.Error != nil {
		return nil, errors.New("Failed to fetch countries")
	}

	// Map models to DTOs
	var responseDTOs []dto.PublicCountryDTO
	for _, country := range countries {
		dto := mapper.CountryToPublicDTO(country)
		responseDTOs = append(responseDTOs, dto)
	}

	// Create paginated response
	response := mapper.CreatePaginatedResponse(responseDTOs, total, page, pageSize)
	return &response, nil
}

func (pc *AuthController) GetDivisions(page, pageSize, countryId int) (*dto.PaginatedResponse, error) {
	// Ensure DB is initialized
	if pc.DB == nil {
		pc.DB = config.DB
	}

	offset := (page - 1) * pageSize

	var country models.Country
	if err := pc.DB.Where("id = ? AND status = ?", countryId, true).First(&country).Error; err != nil {
		return nil, errors.New("Country not found")
	}

	var divisions []models.Division
	var total int64

	pc.DB.Model(&models.Division{}).
		Where("country_id = ? AND status = ?", countryId, true).
		Count(&total)

	result := pc.DB.Where("country_id = ? AND status = ?", countryId, true).
		Preload("Country").
		Order("name ASC").
		Offset(offset).
		Limit(pageSize).
		Find(&divisions)

	if result.Error != nil {
		return nil, errors.New("Failed to fetch divisions")
	}

	var responseDTOs []dto.PublicDivisionDTO
	for _, division := range divisions {
		dto := mapper.DivisionToPublicDTO(division)
		responseDTOs = append(responseDTOs, dto)
	}

	response := mapper.CreatePaginatedResponse(responseDTOs, total, page, pageSize)

	return &response, nil
}

func (pc *AuthController) GetDistrictsByDivision(page, pageSize, divisionId int) (*dto.PaginatedResponse, error) {
	// Ensure DB is initialized
	if pc.DB == nil {
		pc.DB = config.DB
	}

	// Check if division exists and is active
	var division models.Division
	if err := pc.DB.Where("id = ? AND status = ?", divisionId, true).First(&division).Error; err != nil {
		return nil, errors.New("Division not found")
	}
	offset := (page - 1) * pageSize

	var districts []models.District
	var total int64

	// Count total active records for this division
	pc.DB.Model(&models.District{}).
		Where("division_id = ? AND status = ?", divisionId, true).
		Count(&total)

	// Get records with pagination
	result := pc.DB.Where("division_id = ? AND status = ?", divisionId, true).
		Preload("Country").
		Order("name ASC").
		Offset(offset).
		Limit(pageSize).
		Find(&districts)

	if result.Error != nil {
		return nil, errors.New("Failed to fetch districts")
	}

	// Map models to DTOs
	var responseDTOs []dto.PublicDistrictDTO
	for _, district := range districts {
		dto := mapper.DistrictToPublicDTO(district)
		responseDTOs = append(responseDTOs, dto)
	}

	// Create paginated response
	response := mapper.CreatePaginatedResponse(responseDTOs, total, page, pageSize)
	return &response, nil
}
