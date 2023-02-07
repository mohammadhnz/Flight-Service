package middleware

import (
	"awesomeProject/models"
	"awesomeProject/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func RequiredAuth(c *gin.Context) {
	claims, user := ExtractAndDecodeJWTToken(c)
	c.Set("user", user)
	c.Set("token_expiration", claims["exp"])
	c.Next()
}

func ExtractAndDecodeJWTToken(c *gin.Context) (jwt.MapClaims, models.User) {
	tokenString, err := c.Cookie("AuthorizationAccess")

	claims, user, err := utils.ExtractJwtToken(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	return claims, user
}
