package controller

import (
	"awesomeProject/models"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.User_id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString(
		[]byte(os.Getenv("SECRET")),
	)
	return tokenString, err
}
