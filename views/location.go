package views

import (
	"net/http"
	"strconv"

	"github.com/farhapartex/real_estate_be/controllers"
	"github.com/farhapartex/real_estate_be/dto"
	"github.com/gin-gonic/gin"
)

func CreateCountry(ctx *gin.Context, authContoller *controllers.AuthController) {
	var request dto.CountryRequestDTO
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	response, err := authContoller.CreateCountry(request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func CountryList(ctx *gin.Context, authContoller *controllers.AuthController) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10 // Set reasonable limits
	}

	response, total, err := authContoller.ListCountries(page, pageSize)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
		"data":     response,
	})
}

func CountryUpdate(ctx *gin.Context, authContoller *controllers.AuthController) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID"})
		return
	}

	var request dto.CountryUpdateRequestDTO

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	response, err := authContoller.UpdateCountry(uint32(id), request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func CountryDelete(ctx *gin.Context, authContoller *controllers.AuthController) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID"})
		return
	}

	err = authContoller.DeleteCountry(uint32(id))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func CreateDivision(ctx *gin.Context, authContoller *controllers.AuthController) {
	var request dto.DivisionRequestDTO
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	response, err := authContoller.CreateDivision(request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func DivisionList(ctx *gin.Context, authContoller *controllers.AuthController) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10 // Set reasonable limits
	}

	response, total, err := authContoller.DivisionList(page, pageSize)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
		"data":     response,
	})
}

func DivisionUpdate(ctx *gin.Context, authContoller *controllers.AuthController) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid division ID"})
		return
	}

	var request dto.DivisionUpdateRequestDTO

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	response, err := authContoller.UpdateDivision(uint32(id), request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func DivisionDelete(ctx *gin.Context, authContoller *controllers.AuthController) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid division ID"})
		return
	}

	err = authContoller.DeleteDivision(uint32(id))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
