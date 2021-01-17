package seed

import (
	"log"
	"src/src/models"

	"github.com/jinzhu/gorm"
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

	// for i := range admin {
	// 	err = db.Debug().Model(&models.Admin{}).Create(&admin[i]).Error
	// 	if err != nil {
	// 		log.Fatalf("cannot seed users table: %v", err)
	// 	}

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
