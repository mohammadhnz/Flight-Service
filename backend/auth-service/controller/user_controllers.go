package controller

import (
	"awesomeProject/config"
	"awesomeProject/middleware"
	"awesomeProject/models"
	"awesomeProject/repository"
	"awesomeProject/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

func ReturnError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error": message,
	})
}
func SignUp(c *gin.Context) {
	var buser repository.UserData
	if c.Bind(&buser) != nil {
		ReturnError(c, http.StatusBadRequest, "Failed to read body")
		return
	}
	user, err := repository.CreateUser(buser)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, "Failed to create user. Duplicate User")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_data": user,
	})
}

func SignIn(c *gin.Context) {
	var body repository.GetUserData
	if c.Bind(&body) != nil {
		ReturnError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	user, err := repository.GetUser(body)
	if err != nil {
		ReturnError(c, http.StatusForbidden, "Wrong user or password")
		return
	}

	accessTokenString, refreshTokenString, err := utils.GetAuthTokens(user)

	if err != nil {
		ReturnError(c, http.StatusBadRequest, "failed to create token")
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AuthorizationAccess", accessTokenString, 3600, "", "", false, true)
	c.SetCookie("AuthorizationRefresh", refreshTokenString, 12*3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"Yoo": "Yooo"})
}

func SignOut(c *gin.Context) {
	claims, user := middleware.ExtractAndDecodeJWTToken(c)
	token, _ := c.Cookie("Authorization")
	err := SignOutAndUpdateTokens(claims, user, token)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Failed to create",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"ok": "Logged out",
	})

}
func SignOutAndUpdateTokens(claims jwt.MapClaims, user models.User, token string) error {
	if expiration_time, ok := claims["exp"].(float64); ok {
		unauthorizedToken := models.UnauthorizedToken{
			User_id:    user.User_id,
			Token:      token,
			Expiration: time.Unix(int64(expiration_time), 0),
		}
		result := config.DB.Create(&unauthorizedToken)
		if result.Error != nil {
			return errors.New("")
		}
		return nil
	}
	return errors.New("")
}

func UserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"object": user,
	})

}

func All(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"object": users,
	})
}
