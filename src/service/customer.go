package service

import (
	"errors"
	"fmt"
	"html"
	"src/src/db"
	"src/src/models"
	"src/src/util"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type customer struct {
	op         db.CustomerOperation
	kycService KycService
	accService AccountService
}

type CustomerService interface {
	CreateCustomer(c models.Customer) (models.Customer, error)
	Get(customerID string) (models.Customer, error)
	GetAll() ([]models.Customer, error)
	Update(models.Customer) (models.Customer, error)
	Delete(customerID string) error
	LinkKyc(kyc *models.KYCDetails, userID string) (models.Customer, error)
	UnLinkKyc(kyc_id uint32, userID string) (models.Customer, error)
	LinkAccount(acc *models.Account, userID string) (models.Customer, error)
	UnLinkAccount(acc_id uint64, userID string) (models.Customer, error)
}

func NewCustomerService() CustomerService {
	return &customer{
		db.NewCustomerOperation(),
		NewKycService(),
		NewAccountService(),
	}
}
func (e *customer) CreateCustomer(c models.Customer) (models.Customer, error) {
	c.Email = html.EscapeString(strings.TrimSpace(c.Email))
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	encr, err := util.Hash(c.Password)
	if err != nil {
		return models.Customer{}, err
	}
	c.Password = string(encr)
	if err := e.validateCustomer("", c); err != nil {
		return models.Customer{}, err
	}
	return e.op.Create(c)
}

func (e *customer) GetAll() ([]models.Customer, error) {
	return e.op.GetAll()
}
func (e *customer) Get(customerID string) (models.Customer, error) {
	return e.op.Get(customerID)
}

func (e *customer) Update(c models.Customer) (models.Customer, error) {
	c.UpdatedAt = time.Now()
	if c.Email != "" {
		c.Email = html.EscapeString(strings.TrimSpace(c.Email))
	}
	return e.op.Update(c)
}
func (e *customer) Delete(customerID string) error {
	return e.op.Delete(customerID)
}

func (e *customer) LinkKyc(kyc *models.KYCDetails, userID string) (models.Customer, error) {
	kyc.UserID = userID
	k, err := e.kycService.Save(*kyc, userID)
	if err != nil {
		return models.Customer{}, fmt.Errorf("Error in saving KYC details %+v", err)
	}
	emp := models.Customer{
		ID:        userID,
		KycID:     &k.ID,
		UpdatedAt: time.Now(),
	}
	emp, err = e.op.Update(emp)
	if err != nil {
		e := e.kycService.Delete(kyc.ID, userID)
		return models.Customer{}, fmt.Errorf("Error in saving KYC details %+v Reverted back Changes : %+v", err, e)
	}
	kyc = &k
	emp.KycDetails = kyc

	return emp, nil
}
func (e *customer) UnLinkKyc(kycID uint32, userID string) (models.Customer, error) {
	emp, err := e.op.Get(userID)
	if err != nil {
		return models.Customer{}, errors.New("User with the provided ID does not exist")
	}
	if *emp.KycID != kycID {
		return models.Customer{}, errors.New("KYC ID does not match")
	}
	emp.KycID = nil
	emp.UpdatedAt = time.Now()
	return models.Customer{}, e.op.UnlinkKYC(emp)
}
func (e *customer) LinkAccount(acc *models.Account, userID string) (models.Customer, error) {
	acc.UserID = userID
	a, err := e.accService.Save(*acc, userID)
	if err != nil {
		return models.Customer{}, fmt.Errorf("Error in saving acc details %+v", err)
	}
	emp := models.Customer{
		ID:        userID,
		AccountID: &a.ID,
		UpdatedAt: time.Now(),
	}
	emp, err = e.op.Update(emp)
	if err != nil {
		e := e.accService.Delete(acc.ID, userID)
		return models.Customer{}, fmt.Errorf("Error in saving Account details %+v Reverted back Changes : %+v", err, e)
	}
	//acc = &a
	emp.Account = &a
	return emp, nil
}
func (e *customer) UnLinkAccount(acc_id uint64, userID string) (models.Customer, error) {
	emp, err := e.op.Get(userID)
	if err != nil {
		return models.Customer{}, errors.New("User with the provided ID does not exist")
	}
	if *emp.AccountID != acc_id {
		return models.Customer{}, errors.New("Account ID does not match")
	}
	emp.AccountID = nil
	emp.UpdatedAt = time.Now()
	return models.Customer{}, e.op.UnlinkAccount(emp)
}

func (e *customer) validateCustomer(action string, c models.Customer) error {
	switch strings.ToLower(action) {
	case "update":
		if c.Password == "" {
			return errors.New("Required Password")
		}
		if c.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(c.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if c.Password == "" {
			return errors.New("Required Password")
		}
		if c.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(c.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:

		if c.Password == "" {
			return errors.New("Required Password")
		}
		if c.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(c.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}
