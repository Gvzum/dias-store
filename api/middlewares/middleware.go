package middlewares

import (
	"github.com/Gvzum/dias-store.git/api/base"
	"github.com/Gvzum/dias-store.git/config"
	"github.com/Gvzum/dias-store.git/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: token missed",
			})
			return
		}

		_token := strings.TrimPrefix(authHeader, "Bearer ")
		if _token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: token missed",
			})
			return
		}

		token, err := jwt.ParseWithClaims(_token, &base.AuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.Server.TOKEN_SECRET_KEY), nil
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: invalid token",
			})
			return
		}

		claims, ok := token.Claims.(*base.AuthCustomClaims)
		if !ok || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: invalid token",
			})
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}

func ProtectionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, _ := ctx.Get("user_id")
		userIDString, ok := userID.(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}

		user, err := base.GetUserByID(userIDString)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}

func IsSuperUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _ := ctx.Value("user").(*models.User)
		if ctx.Request.Method != http.MethodGet && user.IsSuperUser == false {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "You don't have enough permission",
			})
		}

		ctx.Next()
	}
}
