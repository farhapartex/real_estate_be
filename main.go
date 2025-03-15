package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/farhapartex/real_estate_be/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.MigrateDB()

	r := gin.Default()

	r.GET("/", HealthCheckHandler)
	r.GET("/health_check", HealthCheckHandler)

	err := r.Run(":8080")
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
