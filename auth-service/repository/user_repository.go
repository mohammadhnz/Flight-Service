package repository

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	Email        string
	Phone_number string
	Gender       string
	First_name   string
	Last_name    string
	Password     string
}

type GetUserData struct {
	Email        string
	Phone_number string
	Password     string
}

func CreateUser(userData UserData) (models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)
	if err != nil {
		return models.User{}, err
	}
	user := models.User{
		Email:         userData.Email,
		Phone_number:  userData.Phone_number,
		Gender:        userData.Gender,
		First_name:    userData.First_name,
		Last_name:     userData.Last_name,
		Password_hash: string(hash),
	}
	result := config.DB.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func GetUser(userData GetUserData) (models.User, error) {
	var user models.User
	config.DB.Where("email = ? OR phone_number = ?", userData.Email, userData.Phone_number).First(&user)
	if user.User_id == 0 {
		return models.User{}, errors.New("Wrong user or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(userData.Password))
	if err != nil {
		return models.User{}, errors.New("Wrong user or password")
	}

	return user, nil
}
