package mapper

import (
	"github.com/farhapartex/real_estate_be/dto"
	"github.com/farhapartex/real_estate_be/models"
)

func OwnerSignupDTOToUserModel(request dto.OwnerSignupRequestDTO, hashedPassed string) models.User {
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

func OwnerSignupDTOToProfileModel(request dto.OwnerSignupRequestDTO, userID uint) models.OwnerProfile {

	return models.OwnerProfile{
		UserID:      userID,
		PhoneNumber: request.PhoneNumber,
		CompanyName: nil,
		Website:     nil,
	}
}

func UserToUserDetail(user models.User, profile models.OwnerProfile) dto.UserDetailShortDTO {
	return dto.UserDetailShortDTO{
		ID:            user.ID,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		Role:          string(user.Role),
		IsSuperuser:   user.IsSuperuser,
		JoinedAt:      user.JoinedAt,
		LastLoginAt:   user.LastLoginAt,
		EmailVerified: user.EmailVerified,
		PhoneNumber:   profile.PhoneNumber,
		Website:       profile.Website,
		Status:        user.Status,
	}
}

func UserToLoginResponse(token string) dto.LoginResponseDTO {
	return dto.LoginResponseDTO{
		Token: token,
	}
}

func UserToRegistrationResponse(user models.User, token string) dto.RegisterResponseDTO {
	return dto.RegisterResponseDTO{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      string(user.Role),
		Status:    user.Status,
		Token:     token,
	}
}

func UserToMeResponse(user models.User) dto.UserMeDTO {
	return dto.UserMeDTO{
		ID:            user.ID,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		LastLoginAt:   user.LastLoginAt,
		EmailVerified: user.EmailVerified,
		Role:          string(user.Role),
	}
}

func ToVerifyAccountResponse(success bool, message string) dto.VerifyAccountResponse {
	return dto.VerifyAccountResponse{
		Success: success,
		Message: message,
	}
}

// ToResendVerificationResponse converts resend verification result to a ResendVerificationResponse
func ToResendVerificationResponse(success bool, message string) dto.ResendVerificationResponse {
	return dto.ResendVerificationResponse{
		Success: success,
		Message: message,
	}
}
