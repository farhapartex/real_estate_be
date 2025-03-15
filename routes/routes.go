package routes

import (
	"github.com/farhapartex/real_estate_be/views"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/token", views.SignIn)
		}
	}
}
