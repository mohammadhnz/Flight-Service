package controller

import (
	"awesomeProject/config"
	"awesomeProject/middleware"
	"awesomeProject/models"
	"awesomeProject/repository"
	"github.com/gin-gonic/gin"
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

	tokenString, err := GenerateToken(user)

	if err != nil {
		ReturnError(c, http.StatusBadRequest, "failed to create token")
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"Yoo": "Yooo"})
}

func SignOut(c *gin.Context) {
	claims, user := middleware.DecodeJwtToken(c)
	token, _ := c.Cookie("Authorization")
	if expiration_time, ok := claims["exp"].(float64); ok {
		unauthorizedToken := models.UnauthorizedToken{
			User_id:    user.User_id,
			Token:      token,
			Expiration: time.Unix(int64(expiration_time), 0),
		}
		result := config.DB.Create(&unauthorizedToken)
		if result.Error != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Failed to create",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"ok": "Logged out",
		})
	}

}

func UserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"object": user,
	})

}

func All(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"object": user,
	})
}
