package views

import (
	"net/http"

	"github.com/farhapartex/real_estate_be/controllers"
	"github.com/farhapartex/real_estate_be/dto"
	"github.com/gin-gonic/gin"
)

func SystemAdmin(ctx *gin.Context, authController *controllers.AuthController) {

	err := authController.ConfigureAdmin()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})

}

func SignUp(ctx *gin.Context, authController *controllers.AuthController) {
	var request dto.RegisterRequestDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	response, err := authController.SignUp(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)

}

func SignIn(c *gin.Context, authController *controllers.AuthController) {
	var request dto.LoginRequestDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	response, err := authController.Login(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

}
