package db

import (
	"fmt"
	"log"
	"src/src/models"
	"src/src/seed"

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
	//First Time Activity
	err = db.Debug().DropTableIfExists(&models.Admin{}, &models.Customer{}, &models.Employee{}, &models.Transaction{}, &models.Account{}, &models.KYCDetails{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	db.Debug().AutoMigrate(&models.Admin{}, &models.Customer{}, &models.Employee{}, &models.Account{}, &models.KYCDetails{}, &models.Transaction{}) //database migration
	addConstraintsInCustomerTable()
	addConstraintsInEmployeeTable()
	addConstraintsInAdminTable()
	addConstraintsInTransactionTable()
	seed.Load(db)
}

func addConstraintsInCustomerTable() {
	var err error

	err = db.Debug().Exec("DROP SEQUENCE IF EXISTS customerIDGenerator;").Error
	if err != nil {
		log.Fatalf("Sequence generationg error %v", err)
	}

	err = db.Debug().Exec("CREATE SEQUENCE customerIDGenerator;").Error
	if err != nil {
		log.Fatalf("Sequence generationg error %v", err)
	}
	err = db.Debug().Exec("SELECT setval('customerIDGenerator', 10000);").Error
	if err != nil {
		log.Fatalf("Sequence generationg error %v", err)
	}

	err = db.Debug().Model(&models.Customer{}).AddForeignKey("kyc_id", "kyc_details(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	err = db.Debug().Model(&models.Customer{}).AddForeignKey("account_id", "accounts(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Exec("alter table customers alter column id set default 'sc-' || nextval('customerIDGenerator');").Error
	if err != nil {
		log.Fatalf("Sequence attaching error %v", err)
	}
}

func addConstraintsInEmployeeTable() {
	var err error

	err = db.Debug().Exec("DROP SEQUENCE IF EXISTS employeeIDGenerator;").Error
	if err != nil {
		log.Fatalf("Sequence generationg error %v", err)
	}

	err = db.Debug().Exec("CREATE SEQUENCE employeeIDGenerator;").Error
	if err != nil {
		log.Fatalf("Sequence generationg error %v", err)
	}

	err = db.Debug().Exec("SELECT setval('employeeIDGenerator', 10000);").Error
	if err != nil {
		log.Fatalf("Sequence generationg error %v", err)
	}

	err = db.Debug().Model(&models.Employee{}).AddForeignKey("kyc_id", "kyc_details(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("1 attaching foreign key error: %v", err)
	}
	err = db.Debug().Model(&models.Employee{}).AddForeignKey("account_id", "accounts(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("2 attaching foreign key error: %v", err)
	}

	err = db.Debug().Exec("alter table employees alter column id set default 'se-' || nextval('employeeIDGenerator');").Error
	if err != nil {
		log.Fatalf("Sequence attaching error %v", err)
	}
}

func addConstraintsInAdminTable() {
	var err error
	err = db.Debug().Exec("DROP SEQUENCE IF EXISTS adminIDGenerator;").Error
	if err != nil {
		log.Fatalf("Sequence generationg error %v", err)
	}

	err = db.Debug().Exec("CREATE SEQUENCE adminIDGenerator;").Error
	if err != nil {
		log.Fatalf("Sequence generationg error %v", err)
	}
	err = db.Debug().Exec("SELECT setval('adminIDGenerator', 10000);").Error
	if err != nil {
		log.Fatalf("Sequence generationg error %v", err)
	}
	err = db.Debug().Model(&models.Admin{}).AddForeignKey("kyc_id", "kyc_details(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	err = db.Debug().Model(&models.Admin{}).AddForeignKey("account_id", "accounts(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Exec("alter table admins alter column id set default 'sa-' || nextval('adminIDGenerator');").Error
	if err != nil {
		log.Fatalf("Sequence attaching error %v", err)
	}
}

func addConstraintsInTransactionTable() {
	err := db.Debug().Model(&models.Transaction{}).AddForeignKey("account_id", "accounts(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
}
