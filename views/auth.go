package views

import (
	"fmt"
	"net/http"
	"time"

	"github.com/farhapartex/real_estate_be/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWT(id uint, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &config.Claims{
		Id:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(config.JWTSecret)
}

func SignIn(c *gin.Context) {
	var input struct {
		Id       uint   `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}

	var id uint
	var firstName, lastName, email, strongPassword string

	sql := "SELECT id, first_name, last_name, email, password from users WHERE email = ? ;"
	result := config.DB.Raw(sql, input.Email).Row()

	err = result.Scan(&id, &firstName, &lastName, &email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(strongPassword), []byte(input.Password))
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	token, tokenErr := GenerateJWT(input.Id, input.Email)
	if tokenErr != nil {
		fmt.Println("DEBUG: User not found or scan failed:", tokenErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// update last_login_at
	currentTime := time.Now()
	updateSQL := "UPDATE users SET last_login_at = ? WHERE id = ? ;"
	updateRes := config.DB.Exec(updateSQL, currentTime, id)
	if updateRes.Error != nil {
		fmt.Println("Failed to login")
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
