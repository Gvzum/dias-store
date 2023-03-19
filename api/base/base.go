package base

import (
	"github.com/dgrijalva/jwt-go"
)

type AuthCustomClaims struct {
	UserID string
	jwt.StandardClaims
}
