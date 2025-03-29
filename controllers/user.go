package controllers

import (
	"errors"

	"github.com/farhapartex/real_estate_be/dto"
	"github.com/farhapartex/real_estate_be/mapper"
	"github.com/farhapartex/real_estate_be/models"
	"gorm.io/gorm"
)

func (c *AuthController) GetSystemAllUsers(page, pageSize int) (*dto.PaginatedResponse, error) {
	offset := (page - 1) * pageSize

	var users []models.User
	var total int64

	c.DB.Model(&models.User{}).Count(&total)

	result := c.DB.Order("first_name ASC").Offset(offset).Limit(pageSize).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	var userDTOs []dto.UserDetailShortDTO

	for _, user := range users {
		var profileFound bool = true
		var profile models.OwnerProfile
		c.DB.Where("user_id = ?", user.ID).First(&profile)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				profileFound = false
			} else {
				return nil, result.Error
			}

		}

		if !profileFound {
			profile = models.OwnerProfile{
				UserID:      user.ID,
				CompanyName: "",
				PhoneNumber: "",
				Website:     "",
				CreatedAt:   user.JoinedAt,
				UpdatedAt:   user.JoinedAt,
			}
		}
		userDTOs = append(userDTOs, mapper.UserToUserDetail(user, profile))
	}

	response := mapper.CreatePaginatedResponse(userDTOs, total, page, pageSize)

	return &response, nil
}
