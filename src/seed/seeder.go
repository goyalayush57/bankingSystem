package seed

import (
	"fmt"
	"log"
	"src/src/models"
	"src/src/util"

	"github.com/jinzhu/gorm"
)

func Load(db *gorm.DB) {
	var admin = []models.Admin{
		models.Admin{
			Email:    "ayush@gmail.com",
			Password: "1234",
			Mobile:   "1515",
		},
	}
	// err := db.Debug().DropTableIfExists(&models.Admin{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot drop table: %v", err)
	// }
	// err = db.Debug().AutoMigrate(&models.Admin{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot migrate table: %v", err)
	// }

	for i := range admin {
		encr, _ := util.Hash(admin[i].Password)
		admin[i].Password = string(encr)
		err := db.Debug().Model(&models.Admin{}).Create(&admin[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		fmt.Println("############Admin Details####################")
		fmt.Printf("\n\n%+v\n\n", admin[i])
		fmt.Println("#################################")
	}

	// 	// err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
	// 	// if err != nil {
	// 	// 	log.Fatalf("cannot seed posts table: %v", err)
	// 	// }
	// }

	// res := db.Debug().Model(&models.KYCDetails{}).Create(&models.KYCDetails{
	// 	KYCNumber: "kyc inserted",
	// 	Type:      models.KYCType("Aadhar"),
	// 	UserID:    1,
	// })
	// if res.Error != nil {
	// 	log.Fatalf("cannot seed users table: %v", err)
	// }
	// //var k m
	// var k models.KYCDetails
	// mapstructure.Decode(res.Value, &k)
	// a := models.Admin{}
	// err = db.Debug().Model(models.Admin{}).Where("id = ?", 1).Take(&a).Error
	// if err != nil {
	// 	return
	// }

	// db.Debug().Model(&a).Update("KycDetails", k)
	// b := models.Admin{}
	// err = db.Debug().Model(models.Admin{}).Where("id = ?", 1).Take(&b).Error
	// if err != nil {
	// 	return
	// }
	// fmt.Printf("%+v", b)

}
