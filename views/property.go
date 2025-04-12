package views

import (
	"net/http"
	"strconv"

	"github.com/farhapartex/real_estate_be/controllers"
	"github.com/farhapartex/real_estate_be/dto"
	"github.com/gin-gonic/gin"
)

func PropertieList(ctx *gin.Context, authContoller *controllers.AuthController) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10 // Set reasonable limits
	}

	filters := dto.PropertyFilterDTO{
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
