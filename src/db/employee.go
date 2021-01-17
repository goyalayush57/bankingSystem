package db

import (
	"errors"
	"src/src/models"

	"github.com/jinzhu/gorm"
)

type EmployeeOperation interface {
	Get(employeeID uint32) (models.Employee, error)
	GetAll() ([]models.Employee, error)
	Create(models.Employee) (models.Employee, error)
	Update(models.Employee) error
	Delete(employeeID uint32) error
}

func NewEmployeeService() EmployeeOperation {
	return &employee{}
}

type employee struct{}

func (e *employee) Get(employeeID uint32) (models.Employee, error) {
	emp := models.Employee{}
	err := db.Debug().Model(models.Employee{}).Where("id = ?", employeeID).Take(&emp).Error
	return emp, err
}

func (e *employee) GetAll() ([]models.Employee, error) {
	emps := []models.Employee{}
	err := db.Debug().Model(&models.Employee{}).Find(&emps).Error
	return emps, err
}

func (d *employee) Create(c models.Employee) (models.Employee, error) {
	var err error
	c.Account = models.Account{}
	c.KycDetails = models.KYCDetails{}
	err = db.Debug().Model(&models.Employee{}).Create(&c).Error
	if err != nil {
		return c, err
	}
	return c, nil
}

func (d *employee) Update(c models.Employee) error {
	var err error

	err = db.Debug().Model(&models.Employee{}).Where("id = ?", c.ID).Updates(
		models.Employee{
			KycID:     c.KycID,
			Email:     c.Email,
			AccountID: c.AccountID,
			Mobile:    c.Mobile,
			Password:  c.Password,
		},
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *employee) Delete(employeeID uint32) error {
	db = db.Debug().Model(&models.Employee{}).Where("id = ?", employeeID).Take(&models.Employee{}).Delete(&models.Employee{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return errors.New("employeeID not found")
		}
		return db.Error
	}
	return nil
}
