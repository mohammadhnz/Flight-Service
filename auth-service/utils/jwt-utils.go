package utils

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func GetAuthTokens(user models.User) (string, string, error) {
	accessTokenString, err := GenerateOAuthToken(user, os.Getenv("ACCESS_SECRET"), time.Duration(1*time.Hour))
	refreshTokenStirng, err := GenerateOAuthToken(user, os.Getenv("REFRESH_SECRET"), time.Duration(12*time.Hour))
	return accessTokenString, refreshTokenStirng, err
}

func GenerateOAuthToken(user models.User, secret string, expireTime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.User_id,
		"exp": time.Now().Add(expireTime).Unix(),
	})
	tokenString, err := token.SignedString(
		[]byte(secret),
	)
	return tokenString, err
}

func ExtractJwtToken(tokenString string) (jwt.MapClaims, models.User, error) {
	var claims jwt.MapClaims

	mapClaims, m, err := checkTokenIsNotUnAuthorized(tokenString)
	if err != nil {
		return mapClaims, m, err
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if cl, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims = cl
	} else {
		return nil, models.User{}, errors.New("Unauthorized")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, models.User{}, errors.New("Unauthorized")
	}
	var user models.User
	config.DB.First(&user, claims["sub"])
	if user.User_id == 0 {
		return nil, models.User{}, errors.New("Unauthorized")
	}
	return claims, user, nil
}

func checkTokenIsNotUnAuthorized(tokenString string) (jwt.MapClaims, models.User, error) {
	var unauthorized_toke models.UnauthorizedToken
	config.DB.Where("token = ?", tokenString).First(&unauthorized_toke)
	if unauthorized_toke.User_id != 0 {
		return nil, models.User{}, errors.New("Unauthorized")
	}
	return nil, models.User{}, nil
}
