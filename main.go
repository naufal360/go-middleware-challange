package main

import (
	"go-middleware-challange/database"
	"go-middleware-challange/router"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error .env file at %s", err)
		panic(err)
	}
}

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
