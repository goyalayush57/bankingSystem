package db

import (
	"errors"
	"src/src/models"

	"github.com/jinzhu/gorm"
)

type AccountOperation interface {
	Get(accountID uint32, userID uint32) (models.Account, error)
	Create(acc models.Account, userID uint32) (models.Account, error)
	Delete(accountID uint32, userID uint32) error
	GetAll() ([]models.Account, error)
	Update(acc models.Account) error
}

func NewAccountService() AccountOperation {
	return &account{}
}

type account struct{}

func (a *account) Get(accountID uint32, userID uint32) (models.Account, error) {
	acc := models.Account{}
	err := db.Debug().Model(models.Account{}).Where("id = ? and user_id = ?", accountID, userID).Take(&acc).Error
	return acc, err
}

func (a *account) GetAll() ([]models.Account, error) {
	accs := []models.Account{}
	err := db.Debug().Model(&models.Account{}).Find(&accs).Error
	return accs, err
}

func (a *account) Create(acc models.Account, userID uint32) (models.Account, error) {
	var err error
	err = db.Debug().Model(&models.Account{}).Create(&acc).Error
	if err != nil {
		return acc, err
	}
	return acc, nil
}

func (d *account) Update(acc models.Account) error {
	var err error

	err = db.Debug().Model(&models.Account{}).Where("id = ?", acc.ID).Updates(
		models.Account{
			Balance: acc.Balance,
		},
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *account) Delete(accountID uint32, userID uint32) error {
	db = db.Debug().Model(&models.Account{}).Where("id = ? and user_id = ?", accountID, userID).Take(&models.Account{}).Delete(&models.Account{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return errors.New("accountID not found")
		}
		return db.Error
	}
	return nil
}
