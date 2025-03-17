package routes

import (
	"github.com/farhapartex/real_estate_be/controllers"
	"github.com/farhapartex/real_estate_be/middlewares"
	"github.com/farhapartex/real_estate_be/views"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine, authController *controllers.AuthController) {
	publicApi := r.Group("/api/v1")
	{
		auth := publicApi.Group("/auth")
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

	protectedAPI := r.Group("/api/v1")
	protectedAPI.Use(middlewares.AuthMiddleware())
	{
		protectedAPI.GET("/me", func(ctx *gin.Context) {
			views.Me(ctx, authController)
		})
	}
}
