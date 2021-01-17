package main

import (
	"fmt"
	"log"
	"os"
	"src/src/controller"
	"src/src/db"

	"github.com/joho/godotenv"
)

var database = db.DB{}

func main() {
	fmt.Println("Hello World")
	var err error
	err = godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	database.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	s := &controller.Server{}
	s.Intialize()
	s.Run(":8081")
}
