package database

import (
	"fmt"
	"go-middleware-challange/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	var (
		host   = os.Getenv("DB_HOST")
		user   = os.Getenv("DB_USER")
		dbPort = os.Getenv("DB_PORT")
		dbName = os.Getenv("DB_NAME") // if ur postgres have password, u have add at the dsn variable and at the .env
	)

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable", host, user, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connection to database : %s", err)
	}

	fmt.Println("Success connected to database!")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
