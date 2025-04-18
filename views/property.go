package views

import (
	"net/http"
	"strconv"

	"github.com/farhapartex/real_estate_be/controllers"
	"github.com/farhapartex/real_estate_be/dto"
	"github.com/farhapartex/real_estate_be/models"
	"github.com/gin-gonic/gin"
)

func PropertieList(ctx *gin.Context, authContoller *controllers.AuthController) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID := uint(user.(models.User).ID)

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10 // Set reasonable limits
	}

	filters := dto.PropertyFilterDTO{
		OwerID:  userID,
		Page:    page,
		PerPage: pageSize,
	}

	response, err := authContoller.GetProperties(filters)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func CreateProperty(ctx *gin.Context, authContoller *controllers.AuthController) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID := uint(user.(models.User).ID)

	var request dto.PropertyRequestDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid data", "details": err.Error()})
		return
	}

	response, err := authContoller.CreateProperty(request, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func PropertyDetails(ctx *gin.Context, authContoller *controllers.AuthController) {
	idParam := ctx.Param("id")
	propertyId, err := strconv.ParseUint(idParam, 10, 32)
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID := uint(user.(models.User).ID)

	response, err := authContoller.PropertyDetails(uint32(propertyId), userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func PropertyUpdate(ctx *gin.Context, authContoller *controllers.AuthController) {
	idParam := ctx.Param("id")
	propertyId, err := strconv.ParseUint(idParam, 10, 32)
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID := uint(user.(models.User).ID)
	var request dto.PropertyRequestDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid data", "details": err.Error()})
		return
	}

	response, err := authContoller.PropertyPatch(uint32(propertyId), userID, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
