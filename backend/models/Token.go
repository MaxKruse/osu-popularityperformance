package models

import "golang.org/x/oauth2"

type Token struct {
	JsonModel
	UserId       uint   `gorm:"not null"`
	SessionToken string `gorm:"not null"`
	oauth2.Token
}
