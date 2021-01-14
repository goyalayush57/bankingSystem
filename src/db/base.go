package db

import (
	"fmt"
	"log"
	"src/src/models"
	"src/src/seed"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type DB struct {
	DB *gorm.DB
}

func (server *DB) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.Admin{}, &models.Customer{}, &models.Employee{}, &models.Account{}, &models.KYCDetails{}, &models.Transaction{}) //database migration
	seed.Load(server.DB)
}

func (server *DB) Run(addr string) {
	fmt.Println("Listening to port 8080")
	//log.Fatal(http.ListenAndServe(addr, server.Router))
}
