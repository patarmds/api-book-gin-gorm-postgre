package database

import (
	"fmt"
	"rest-api-books-gin-gorm/models"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var (
	host = "localhost"
	port = "5432"
	user = ""
	password = ""
	dbname = ""
	db *gorm.DB
	err error
)

func StartDB()  {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",host, user, password, dbname, port)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to databse:", err)
	}

	db.Debug().AutoMigrate(models.Book{})
}

func GetDB() *gorm.DB{
	return db
}