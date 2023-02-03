package controller

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignUp(c *gin.Context) {
	//read body
	var buser struct {
		email        string
		phone_number string
		gender       string
		first_name   string
		last_name    string
		password     string
	}
	if c.BindJSON(&buser) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(buser.password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body" + err.Error(),
		})
		return
	}
	user := models.User{
		Email:         buser.email,
		Phone_number:  buser.phone_number,
		Gender:        buser.gender,
		First_name:    buser.first_name,
		Last_name:     buser.last_name,
		Password_hash: string(hash),
	}
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create",
		})
		return
	}
	c.String(200, "Hello world")
}

func SignIn(c *gin.Context) {
	c.String(200, "Hello world")
}

func SignOut(c *gin.Context) {
	c.String(200, "Hello world")
}

func UserInfo(c *gin.Context) {
	c.String(200, "Hello world")
}
