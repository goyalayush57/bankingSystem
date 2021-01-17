package db

import (
	"errors"
	"src/src/models"

	"github.com/jinzhu/gorm"
)

type KycOperation interface {
	Get(userID uint32) (models.KYCDetails, error)
	GetFromID(kycID uint32) (models.KYCDetails, error)
	Update(models.KYCDetails) error
	Delete(kycID uint32, userID uint32) error
	GetAll() ([]models.KYCDetails, error)
}
type kyc struct{}

func NewKycOperation() KycOperation {
	return &kyc{}
}

func (k *kyc) GetFromID(kycID uint32) (models.KYCDetails, error) {
	kyc := models.KYCDetails{}
	err := db.Debug().Model(models.KYCDetails{}).Where("id = ?", kycID).Take(&kyc).Error
	return kyc, err
}

func (k *kyc) Get(userID uint32) (models.KYCDetails, error) {
	kyc := models.KYCDetails{}
	err := db.Debug().Model(models.KYCDetails{}).Where("user_id = ?", userID).Take(&kyc).Error
	return kyc, err
}
func (k *kyc) Update(kyc models.KYCDetails) error {
	var err error

	err = db.Debug().Model(&models.KYCDetails{}).Where("id = ? and user_id = ?", kyc.ID, kyc.UserID).Updates(
		models.KYCDetails{
			KYCNumber: kyc.KYCNumber,
			Type:      kyc.Type,
		},
	).Error
	if err != nil {
		return err
	}
	return nil
}
func (k *kyc) Delete(kycID uint32, userID uint32) error {
	db = db.Debug().Model(&models.KYCDetails{}).Where("id = ? and user_id = ?", kycID, userID).Take(&models.KYCDetails{}).Delete(&models.KYCDetails{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return errors.New("accountID not found")
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
