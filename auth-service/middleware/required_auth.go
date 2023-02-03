package middleware

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"time"
)

func RequiredAuth(c *gin.Context) {
	claims, user := DecodeJwtToken(c)
	c.Set("user", user)
	c.Set("token_expiration", claims["exp"])
	c.Next()
}

func DecodeJwtToken(c *gin.Context) (jwt.MapClaims, models.User) {
	var claims jwt.MapClaims
	tokenString, err := c.Cookie("Authorization")
	var unauthorized_toke models.UnauthorizedToken
	config.DB.Where("token = ?", tokenString).First(&unauthorized_toke)
	if unauthorized_toke.User_id != 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if cl, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims = cl
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	var user models.User
	config.DB.First(&user, claims["sub"])

	if user.User_id == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	return claims, user
}
