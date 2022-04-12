package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	BanchoUser
	Session []Session
}
