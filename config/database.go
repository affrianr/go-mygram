package config

import (
	"fmt"

	"github.com/affrianr/go-mygram/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, user, password, dbname)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&entity.User{}, &entity.SocialMedia{}, &entity.Photo{}, &entity.Comment{})
}
