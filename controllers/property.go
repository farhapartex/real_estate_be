package controllers

import (
	"errors"
	"fmt"

	"github.com/farhapartex/real_estate_be/dto"
	"github.com/farhapartex/real_estate_be/mapper"
	"github.com/farhapartex/real_estate_be/models"
)

func (c *AuthController) GetProperties(filter dto.PropertyFilterDTO) (*dto.PaginatedResponse, error) {
	var properties []models.Property
	var total int64

	// Build query with filters
	query := c.DB.Model(&models.Property{})

	query = query.Where("owner_id = ?", filter.OwerID)

	// Apply filters
	if filter.Purpose != "" {
		query = query.Where("purpose = ?", filter.Purpose)
	}

	if filter.MinPrice > 0 {
		query = query.Where("price >= ?", filter.MinPrice)
	}

	if filter.MaxPrice > 0 {
		query = query.Where("price <= ?", filter.MaxPrice)
	}

	if filter.PropertyType != "" {
		query = query.Where("property_type = ?", filter.PropertyType)
	}

	if filter.BedRooms > 0 {
		query = query.Where("bed_rooms >= ?", filter.BedRooms)
	}

	if filter.BathRooms > 0 {
		query = query.Where("bath_rooms >= ?", filter.BathRooms)
	}

	if filter.MinSize > 0 {
		query = query.Where("size >= ?", filter.MinSize)
	}

	if filter.MaxSize > 0 {
		query = query.Where("size <= ?", filter.MaxSize)
	}

	if filter.CountryID > 0 {
		query = query.Where("country_id = ?", filter.CountryID)
	}

	if filter.DivisionID > 0 {
		query = query.Where("division_id = ?", filter.DivisionID)
	}

	if filter.DistrictID > 0 {
		query = query.Where("district_id = ?", filter.DistrictID)
	}

	if filter.Status != "" {
		fmt.Println("Status filter applied:", filter.Status)
		query = query.Where("status = ?", filter.Status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, errors.New("error counting properties")
	}

	offset := (filter.Page - 1) * filter.PerPage

	// Execute query with pagination
	err := query.Preload("Country").
		Preload("Division").
		Preload("District").
		Preload("Owner").
		Offset(offset).
		Limit(filter.PerPage).
		Order("created_at DESC").
		Find(&properties).Error

	if err != nil {
		return nil, errors.New("error retrieving properties")
	}

	var responseDTOs []dto.PropertyListDTO
	for _, property := range properties {
		fmt.Print("Status:", property.Status)
		dto := mapper.PropertyModelToResponseDTOMapper(property)
		responseDTOs = append(responseDTOs, dto)
	}

	response := mapper.CreatePaginatedResponse(responseDTOs, total, filter.Page, filter.PerPage)

	return &response, nil
}

func (c *AuthController) CreateProperty(request dto.PropertyRequestDTO, userID uint) (*dto.PropertyListDTO, error) {
	// Verify that country, division, and district exist
	var country models.Country
	var division models.Division
	var district models.District

	if err := c.DB.First(&country, request.CountryID).Error; err != nil {
		return nil, errors.New("country not found")
	}

	if err := c.DB.First(&division, request.DivisionID).Error; err != nil {
		return nil, errors.New("division not found")
	}

	if err := c.DB.First(&district, request.DistrictID).Error; err != nil {
		return nil, errors.New("district not found")
	}

	// Create new property
	newProperty := mapper.PropertyDtoToModelMapper(request, userID)

	tx := c.DB.Begin()
	if err := tx.Create(&newProperty).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("property creation failed: " + err.Error())
	}

	// Fetch the created property with all relationships to build the response
	if err := tx.Preload("Country").Preload("Division").Preload("District").Preload("Owner").First(&newProperty, newProperty.ID).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("error retrieving created property")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errors.New("property creation failed during commit")
	}

	response := mapper.PropertyModelToResponseDTOMapper(newProperty)
	return &response, nil
}
