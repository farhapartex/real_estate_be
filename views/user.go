package views

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/farhapartex/real_estate_be/controllers"
	"github.com/farhapartex/real_estate_be/dto"
	"github.com/gin-gonic/gin"
)

func SystemAllUserListView(ctx *gin.Context, authController *controllers.AuthController) {
	page, pageSize := GetPaginationParams(ctx)

	var filter dto.UserFilterDTO

	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid filter parameters",
			"details": err.Error(),
		})
		return
	}

	emailVerifiedStr := ctx.Query("email_verified")
	fmt.Println("email_verified:", emailVerifiedStr)
	if emailVerifiedStr != "" {
		// Convert to boolean based on string value
		if emailVerifiedStr == "true" || emailVerifiedStr == "1" {
			trueValue := true
			filter.EmailVerified = &trueValue
		} else if emailVerifiedStr == "false" || emailVerifiedStr == "0" {
			falseValue := false
			filter.EmailVerified = &falseValue
		} else {
			// If value is invalid, ignore this filter
			filter.EmailVerified = nil
		}
	} else if ctx.Request.URL.RawQuery != "" &&
		strings.Contains(ctx.Request.URL.RawQuery, "email_verified") {
		filter.EmailVerified = nil
	}

	response, err := authController.GetSystemAllUsers(page, pageSize, filter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)

}
