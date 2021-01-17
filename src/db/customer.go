package db

import (
	"errors"
	"src/src/models"

	"github.com/jinzhu/gorm"
)

type CustomerOperation interface {
	Get(customerID string) (models.Customer, error)
	GetAll() ([]models.Customer, error)
	Create(models.Customer) (models.Customer, error)
	Update(models.Customer) (models.Customer, error)
	Delete(customerID string) error
	UnlinkKYC(customer models.Customer) error
	UnlinkAccount(customer models.Customer) error
}

func NewCustomerOperation() CustomerOperation {
	return &customer{}
}

type customer struct{}

func (d *customer) Get(customerID string) (models.Customer, error) {
	c := models.Customer{}
	err := db.Debug().Model(models.Customer{}).Where("id = ?", customerID).Take(&c).Error
	return c, err
}

func (d *customer) GetAll() ([]models.Customer, error) {
	customers := []models.Customer{}
	err := db.Debug().Model(&models.Customer{}).Find(&customers).Error
	return customers, err
}

func (d *customer) Create(c models.Customer) (models.Customer, error) {
	var err error
	//c.Account = &models.Account{}
	//	c.KycDetails = &models.KYCDetails{}
	err = db.Debug().Model(&models.Customer{}).Create(&c).Error
	if err != nil {
		return c, err
	}
	// if c.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
	// 	if err != nil {
	// 		return &Post{}, err
	// 	}
	// }
	return c, nil
}

func (d *customer) Update(c models.Customer) (models.Customer, error) {
	var err error

	err = db.Debug().Model(&models.Customer{}).Where("id = ?", c.ID).Updates(
		models.Customer{
			KycID:     c.KycID,
			Email:     c.Email,
			AccountID: c.AccountID,
			Mobile:    c.Mobile,
			Password:  c.Password,
			UpdatedAt: c.UpdatedAt,
		},
	).Error
	if err != nil {
		return models.Customer{}, err
	}
	// This is the display the updated user
	err = db.Debug().Model(&models.Customer{}).Where("id = ?", c.ID).Take(&c).Error
	if err != nil {
		return models.Customer{}, err
	}
	return c, nil
}

func (d *customer) Delete(customerID string) error {
	db = db.Debug().Model(&models.Customer{}).Where("id = ?", customerID).Take(&models.Customer{}).Delete(&models.Customer{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return errors.New("Customer not found")
		}
		return db.Error
	}
	return nil
}

func (d *customer) UnlinkKYC(cust models.Customer) error {
	return db.Debug().Exec("Update customers SET kyc_id = NULL where id = ?", cust.ID).Error
}

func (d *customer) UnlinkAccount(cust models.Customer) error {
	return db.Debug().Exec("Update customers SET account_id = NULL where id = ?", cust.ID).Error
}
