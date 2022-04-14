package models

import (
	"time"

	"gorm.io/gorm"
)

type JsonModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	JsonModel
	BanchoId    int       `json:"user_id"`
	Username    string    `json:"username"`
	Sessions    []Session `json:"-"`
	OAuthTokens *Token    `json:"-"`
}
