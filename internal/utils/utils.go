package utils

import (
	"github.com/Gvzum/dias-store.git/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(username string) (string, error) {
	// Create a new token object, specifying signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	})

	// Sign the token with a secret key
	tokenSecret := config.AppConfig.Server.TOKEN_SECRET_KEY
	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
