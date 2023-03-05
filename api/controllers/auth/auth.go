package auth

import (
	"fmt"
	"github.com/Gvzum/dias-store.git/internal/models"
	"github.com/Gvzum/dias-store.git/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type Controller struct{}

var userModel = new(models.User)

func (u Controller) Create(c *gin.Context) {
	var user UserAuthorization
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Print(user)

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}

func (u Controller) Login(c *gin.Context) {
	var user UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(user)
}

// registerHandler is a handler for user registration
func registerHandler(c *gin.Context) {
	// Parse request body
	var user UserAuthorization
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	user.Password = string(hashedPassword)
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate JWT token
	tokenString, err := utils.GenerateToken(user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	// Return token
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
