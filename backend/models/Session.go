package models

type Session struct {
	JsonModel
	UserId       uint   `gorm:"not null"`
	SessionToken string `gorm:"not null"`
}
