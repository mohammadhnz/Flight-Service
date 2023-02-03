package controller

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func SignUp(c *gin.Context) {
	var buser struct {
		Email        string
		Phone_number string
		Gender       string
		First_name   string
		Last_name    string
		Password     string
	}
	if c.Bind(&buser) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(buser.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body" + err.Error(),
		})
		return
	}
	user := models.User{
		Email:         buser.Email,
		Phone_number:  buser.Phone_number,
		Gender:        buser.Gender,
		First_name:    buser.First_name,
		Last_name:     buser.Last_name,
		Password_hash: string(hash),
	}
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Failed to create",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"ok": "user created",
	})
}

func SignIn(c *gin.Context) {
	//validate user data
	var body struct {
		Email        string
		Phone_number string
		Password     string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	var user models.User
	config.DB.Where("email = ? OR phone_number = ?", body.Email, body.Phone_number).First(&user)
	if user.User_id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no user with given information"})
	}

	//validate password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no user with given information",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.User_id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString(
		[]byte(os.Getenv("SECRET")),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func SignOut(c *gin.Context) {
	c.String(200, "Hello world")
}

func UserInfo(c *gin.Context) {
	c.String(200, "Hello world")
}

func All(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}
