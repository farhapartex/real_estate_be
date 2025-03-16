package mapper

import (
	"github.com/farhapartex/real_estate_be/dto"
	"github.com/farhapartex/real_estate_be/models"
)

func RegisterRequestToUserModel(request dto.RegisterRequestDTO, hashedPassed string) models.User {
	return models.User{
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Email:         request.Email,
		Password:      hashedPassed,
		IsSuperuser:   false,
		Role:          models.Role(models.OwnerRole),
		Status:        "inactive",
		EmailVerified: false,
	}
}

func UserToUserDetail(user models.User) dto.UserDetailShortDTO {
	return dto.UserDetailShortDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func UserToLoginResponse(user models.User, token string) dto.LoginResponseDTO {
	return dto.LoginResponseDTO{
		Token: token,
		User:  UserToUserDetail(user),
	}
}

func UserToRegistrationResponse(user models.User) dto.RegisterResponseDTO {
	return dto.RegisterResponseDTO{
		FirstName: user.FirstName,
	}
}
