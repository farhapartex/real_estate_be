package routes

import (
	"github.com/farhapartex/real_estate_be/controllers"
	"github.com/farhapartex/real_estate_be/views"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine, authController *controllers.AuthController) {
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/token", func(ctx *gin.Context) {
				views.SignIn(ctx, authController)
			})

			auth.POST("/signup", func(ctx *gin.Context) {
				views.SignUp(ctx, authController)
			})

			auth.POST("/admin", func(ctx *gin.Context) {
				views.SystemAdmin(ctx, authController)
			})
		}
	}
}
