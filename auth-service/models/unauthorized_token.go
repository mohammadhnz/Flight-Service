package models

import "time"

type UnauthorizedToken struct {
	User_id    int64
	Token      string
	Expiration time.Time
}

func (UnauthorizedToken) TableName() string {
	return "unauthorized_token"
}
