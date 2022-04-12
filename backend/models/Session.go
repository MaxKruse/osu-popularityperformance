package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	UserId       uint   `gorm:"not null"`
	SessionToken string `gorm:"not null"`
}
