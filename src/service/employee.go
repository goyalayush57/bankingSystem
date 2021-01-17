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

type employee struct {
	op         db.EmployeeOperation
	kycService KycService
	accService AccountService
}

type EmployeeService interface {
	CreateEmployee(emp models.Employee) (models.Employee, error)
	Get(employeeID string) (models.Employee, error)
	GetAll() ([]models.Employee, error)
	Update(models.Employee) (models.Employee, error)
	Delete(employeeID string) error
	LinkKyc(kyc models.KYCDetails, userID string) (models.Employee, error)
	UnLinkKyc(kycID uint32, userID string) (models.Employee, error)
	LinkAccount(acc models.Account, userID string) (models.Employee, error)
	UnLinkAccount(accID uint64, userID string) (models.Employee, error)
}

func NewEmployeeService() EmployeeService {
	return &employee{
		db.NewEmployeeOperation(),
		NewKycService(),
		NewAccountService(),
	}
}
func (e *employee) CreateEmployee(emp models.Employee) (models.Employee, error) {

	emp.Email = html.EscapeString(strings.TrimSpace(emp.Email))
	emp.CreatedAt = time.Now()
	emp.UpdatedAt = time.Now()

	encr, err := util.Hash(emp.Password)
	if err != nil {
		return models.Employee{}, err
	}
	emp.Password = string(encr)
	if err := e.validateEmployee("", emp); err != nil {
		return emp, err
	}
	//Validate
	return e.op.Create(emp)
}

func (e *employee) GetAll() ([]models.Employee, error) {
	return e.op.GetAll()
}
func (e *employee) Get(employeeID string) (models.Employee, error) {
	return e.op.Get(employeeID)
}

func (e *employee) Update(emp models.Employee) (models.Employee, error) {
	emp.UpdatedAt = time.Now()
	if emp.Email != "" {
		emp.Email = html.EscapeString(strings.TrimSpace(emp.Email))
	}
	return e.op.Update(emp)
}
func (e *employee) Delete(employeeID string) error {
	return e.op.Delete(employeeID)
}

func (e *employee) LinkKyc(kyc models.KYCDetails, userID string) (models.Employee, error) {
	kyc, err := e.kycService.Save(kyc, userID)
	if err != nil {
		return models.Employee{}, fmt.Errorf("Error in saving KYC details %+v", err)
	}
	emp := models.Employee{
		ID:        userID,
		KycID:     &kyc.ID,
		UpdatedAt: time.Now(),
	}
	emp, err = e.op.Update(emp)
	if err != nil {
		e := e.kycService.Delete(kyc.ID, userID)
		return models.Employee{}, fmt.Errorf("Error in saving KYC details %+v Reverted back Changes : %+v", err, e)
	}
	return emp, nil
}
func (e *employee) UnLinkKyc(kycID uint32, userID string) (models.Employee, error) {
	emp, err := e.op.Get(userID)
	if err != nil {
		return models.Employee{}, errors.New("User with the provided ID does not exist")
	}
	if emp.KycID != &kycID {
		return models.Employee{}, errors.New("KYC ID does not match")
	}
	emp.KycID = nil
	emp.UpdatedAt = time.Now()
	return e.op.Update(emp)
}
func (e *employee) LinkAccount(acc models.Account, userID string) (models.Employee, error) {
	acc, err := e.accService.Save(acc, userID)
	if err != nil {
		return models.Employee{}, fmt.Errorf("Error in saving acc details %+v", err)
	}
	emp := models.Employee{
		ID:        userID,
		AccountID: &acc.ID,
		UpdatedAt: time.Now(),
	}
	emp, err = e.op.Update(emp)
	if err != nil {
		e := e.accService.Delete(acc.ID, userID)
		return models.Employee{}, fmt.Errorf("Error in saving Account details %+v Reverted back Changes : %+v", err, e)
	}
	return emp, nil
}
func (e *employee) UnLinkAccount(acc_id uint64, userID string) (models.Employee, error) {
	emp, err := e.op.Get(userID)
	if err != nil {
		return models.Employee{}, errors.New("User with the provided ID does not exist")
	}
	if emp.AccountID != &acc_id {
		return models.Employee{}, errors.New("Account ID does not match")
	}
	emp.AccountID = nil
	emp.UpdatedAt = time.Now()
	return e.op.Update(emp)
}

func (e *employee) validateEmployee(action string, emp models.Employee) error {
	switch strings.ToLower(action) {
	case "update":
		if emp.Password == "" {
			return errors.New("Required Password")
		}
		if emp.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(emp.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if emp.Password == "" {
			return errors.New("Required Password")
		}
		if emp.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(emp.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:

		if emp.Password == "" {
			return errors.New("Required Password")
		}
		if emp.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(emp.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}
