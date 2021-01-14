package seed

import (
	"fmt"
	"log"
	"src/src/models"

	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
)

var admin = []models.Admin{
	models.Admin{
		Email:    "steven@gmail.com",
		Password: "password",
		Mobile:   "1515",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Admin{}, &models.KYCDetails{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Admin{}, &models.KYCDetails{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	for i := range admin {
		err = db.Debug().Model(&models.Admin{}).Create(&admin[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		// err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		// if err != nil {
		// 	log.Fatalf("cannot seed posts table: %v", err)
		// }
	}
	res := db.Debug().Model(&models.KYCDetails{}).Create(&models.KYCDetails{
		KYCNumber: "kyc inserted",
		Type:      models.KYCType("Aadhar"),
	})
	if res.Error != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	//var k m
	var k models.KYCDetails
	mapstructure.Decode(res.Value, &k)
	fmt.Println(k.ID)
	var a models.Admin
	db.Debug().First(&a, 1)
	fmt.Printf("%+v", a)
	db.Debug().Model(&a).Update("kyc_id", k.ID)
}
