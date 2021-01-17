package db

import (
	"fmt"
	"log"
	"src/src/models"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type DB struct{}

var db *gorm.DB

func (server *DB) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		db, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	//to be removed
	// err = db.Debug().DropTableIfExists(&models.Admin{}, &models.Customer{}, &models.Employee{}, &models.Account{}, &models.KYCDetails{}, &models.Transaction{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot drop table: %v", err)
	// }
	addConstraintsInCustomerTable()
	addConstraintsInEmployeeTable()
	db.Debug().AutoMigrate(&models.Admin{}, &models.Customer{}, &models.Employee{}, &models.Account{}, &models.KYCDetails{}, &models.Transaction{}) //database migration

	//seed.Load(db)
}

func addConstraintsInCustomerTable() {
	err := db.Debug().Model(&models.Customer{}).AddForeignKey("kyc_id", "kyc_details(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	err = db.Debug().Model(&models.Customer{}).AddForeignKey("account_id", "accounts(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
}

func addConstraintsInEmployeeTable() {
	err := db.Debug().Model(&models.Employee{}).AddForeignKey("kyc_id", "kyc_details(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	err = db.Debug().Model(&models.Employee{}).AddForeignKey("account_id", "accounts(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
}

func (server *DB) Run(addr string) {
	fmt.Println("Listening to port 8080")
	//log.Fatal(http.ListenAndServe(addr, server.Router))
}
