package models

type User struct {
	user_id       int64
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
