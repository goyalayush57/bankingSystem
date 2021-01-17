package db

import (
	"errors"

	"src/src/models"

	"github.com/jinzhu/gorm"
)

type EmployeeOperation interface {
	Get(employeeID string) (models.Employee, error)
	GetAll() ([]models.Employee, error)
	Create(models.Employee) (models.Employee, error)
	Update(models.Employee) (models.Employee, error)
	Delete(employeeID string) error
}

func NewEmployeeOperation() EmployeeOperation {
	return &employee{}
}

type employee struct{}

func (e *employee) Get(employeeID string) (models.Employee, error) {
	emp := models.Employee{}
	err := db.Debug().Model(models.Employee{}).Where("id = ?", employeeID).Take(&emp).Error
	return emp, err
}

func (e *employee) GetAll() ([]models.Employee, error) {
	emps := []models.Employee{}
	err := db.Debug().Model(&models.Employee{}).Find(&emps).Error
	return emps, err
}

func (e *employee) Create(c models.Employee) (models.Employee, error) {
	var err error
	//c.Account = &models.Account{}
	//c.KycDetails = &models.KYCDetails{}
	err = db.Debug().Model(&models.Employee{}).Create(&c).Error
	if err != nil {
		return c, err
	}
	return c, nil
}

func (e *employee) Update(c models.Employee) (models.Employee, error) {
	var err error

	err = db.Debug().Model(&models.Employee{}).Where("id = ?", c.ID).Updates(
		models.Employee{
			KycID:     c.KycID,
			Email:     c.Email,
			AccountID: c.AccountID,
			Mobile:    c.Mobile,
			Password:  c.Password,
			UpdatedAt: c.UpdatedAt,
		},
	).Error

	if err != nil {
		return models.Employee{}, err
	}
	// This is the display the updated user
	err = db.Debug().Model(&models.Employee{}).Where("id = ?", c.ID).Take(&c).Error
	if err != nil {
		return models.Employee{}, err
	}
	return c, nil
}

func (e *employee) Delete(employeeID string) error {
	db = db.Debug().Model(&models.Employee{}).Where("id = ?", employeeID).Take(&models.Employee{}).Delete(&models.Employee{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return errors.New("employeeID not found")
		}
		return db.Error
	}
	return nil
}
