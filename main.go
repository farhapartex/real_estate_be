package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/farhapartex/real_estate_be/config"
	"github.com/farhapartex/real_estate_be/controllers"
	"github.com/farhapartex/real_estate_be/middlewares"
	"github.com/farhapartex/real_estate_be/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.MigrateDB()

	authController := controllers.NewAuthController(config.DB)

	r := gin.Default()

	// setup middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())

	// setup routes
	routes.RegisterRoute(r, authController)

	// system health check route
	r.GET("/", HealthCheckHandler)
	r.GET("/health_check", HealthCheckHandler)

	port := os.Getenv("port")

	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("Error from main: ", err)
	}

	fmt.Println("server running on port 8080")
}

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "System is up and running",
	})
}
