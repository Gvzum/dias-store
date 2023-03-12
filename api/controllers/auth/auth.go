package auth

import (
	"github.com/Gvzum/dias-store.git/api"
	"github.com/Gvzum/dias-store.git/config"
	"github.com/Gvzum/dias-store.git/config/database"
	"github.com/Gvzum/dias-store.git/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type authCustomClaims struct {
	UserID string
	jwt.StandardClaims
}

type Controller struct{}

func (c Controller) SignInHandler(ctx *gin.Context) {
	var validatedUser UserSignIn
	if err := ctx.ShouldBindJSON(&validatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Get user by email
	user, err := api.GetUserByEmail(validatedUser.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(validatedUser.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password.",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &authCustomClaims{
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	// sign token with secret key and generate signed string
	signedToken, err := token.SignedString([]byte(config.AppConfig.Server.TOKEN_SECRET_KEY))
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"token": signedToken})

}

func (c Controller) SignUpHandler(ctx *gin.Context) {
	var validatedUser UserSignUp
	if err := ctx.ShouldBindJSON(&validatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(validatedUser.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{Name: validatedUser.Name, Email: validatedUser.Email, Password: string(hashedPassword)}

	// Create user
	db := database.GetDB()
	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})

}
