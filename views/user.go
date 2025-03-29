package views

import (
	"net/http"

	"github.com/farhapartex/real_estate_be/controllers"
	"github.com/gin-gonic/gin"
)

func SystemAllUserListView(ctx *gin.Context, authController *controllers.AuthController) {
	page, pageSize := GetPaginationParams(ctx)

	response, err := authController.GetSystemAllUsers(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)

}
