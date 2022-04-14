package database

import (
	"fmt"
	"os"

	"github.com/maxkruse/osu-popularityperformance/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	// make database connection
	host := "database"

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, "5432", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate structs
	db.AutoMigrate(&models.Token{})
	db.AutoMigrate(&models.Session{})
	db.AutoMigrate(&models.User{})

}

func GetSession() *gorm.DB {
	return db.Session(&gorm.Session{})
}
