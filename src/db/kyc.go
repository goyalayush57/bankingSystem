package db

import (
	"errors"
	"fmt"
	"src/src/models"

	"github.com/jinzhu/gorm"
)

type KycOperation interface {
	Create(kyc models.KYCDetails) (models.KYCDetails, error)
	Get(userID string) (models.KYCDetails, error)
	GetFromID(kycID uint32) (models.KYCDetails, error)
	Update(models.KYCDetails) (models.KYCDetails, error)
	Delete(kycID uint32, userID string) error
	GetAll() ([]models.KYCDetails, error)
}
type kyc struct{}

func NewKycOperation() KycOperation {
	return &kyc{}
}
func (k *kyc) Create(kyc models.KYCDetails) (models.KYCDetails, error) {
	var err error
	err = db.Debug().Model(&models.KYCDetails{}).Create(&kyc).Error
	fmt.Println(kyc)
	if err != nil {
		return models.KYCDetails{}, err
	}
	return kyc, nil
}

func (k *kyc) GetFromID(kycID uint32) (models.KYCDetails, error) {
	kyc := models.KYCDetails{}
	err := db.Debug().Model(models.KYCDetails{}).Where("id = ?", kycID).Take(&kyc).Error
	return kyc, err
}

func (k *kyc) Get(userID string) (models.KYCDetails, error) {
	kyc := models.KYCDetails{}
	err := db.Debug().Model(models.KYCDetails{}).Where("user_id = ?", userID).Take(&kyc).Error
	return kyc, err
}
func (k *kyc) Update(kyc models.KYCDetails) (models.KYCDetails, error) {
	var err error

	err = db.Debug().Model(&models.KYCDetails{}).Where("id = ? and user_id = ?", kyc.ID, kyc.UserID).Updates(
		models.KYCDetails{
			KYCNumber: kyc.KYCNumber,
			Type:      kyc.Type,
		},
	).Error
	if err != nil {
		return models.KYCDetails{}, err
	}
	// This is the display the updated kyc
	err = db.Debug().Model(&models.KYCDetails{}).Where("id = ?", kyc.ID).Take(&kyc).Error
	if err != nil {
		return models.KYCDetails{}, err
	}
	return kyc, nil
}
func (k *kyc) Delete(kycID uint32, userID string) error {
	db = db.Debug().Model(&models.KYCDetails{}).Where("id = ? and user_id = ?", kycID, userID).Take(&models.KYCDetails{}).Delete(&models.KYCDetails{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return errors.New("kycID not found")
		}
		return db.Error
	}
	return nil
}
func (k *kyc) GetAll() ([]models.KYCDetails, error) {
	kycs := []models.KYCDetails{}
	err := db.Debug().Model(&models.Account{}).Find(&kycs).Error
	return kycs, err
}
