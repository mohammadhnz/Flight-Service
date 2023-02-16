package models

type User struct {
	User_id       int64 `gorm:"autoIncrement"`
	Email         string
	Phone_number  string
	Gender        string
	First_name    string
	Last_name     string
	Password_hash string
}

func (User) TableName() string {
	return "user_account"
}
