package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type User struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}
type Admin struct {
	ID         string      `sql:"primary_key;default 'SE'||nextval('adminIDGenerator')"  json:"id"`
	Name       string      `gorm:"size:100;not null" json:"name"`
	CreatedAt  time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	KycID      *uint32     `json:"kyc_id"`
	KycDetails *KYCDetails `json:"kyc,omitempty"`
	AccountID  *uint64     `json:"account_id"`
	Account    *Account    `json:"account,omitempty"`
	Mobile     string      `gorm:"size:100;not null;unique" json:"mobile"`
	Email      string      `gorm:"size:100;not null;unique" json:"email"`
	Password   string      `gorm:"size:100;not null;" json:"-"`
	UpdatedAt  time.Time
}
type Employee struct {
	ID         string      `sql:"primary_key;default 'SE'||nextval('employeeIDGenerator')" json:"id"`
	Name       string      `gorm:"size:100;not null" json:"name"`
	CreatedAt  time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	KycID      *uint32     `json:"kyc_id,default:-1"`
	KycDetails *KYCDetails `json:"kyc,omitempty"`
	AccountID  *uint64     `json:"account_id"`
	Account    *Account    `json:"account,omitempty"`
	Mobile     string      `gorm:"size:100;not null;unique" json:"mobile"`
	Email      string      `gorm:"size:100;not null;unique" json:"email"`
	Password   string      `gorm:"size:100;not null;" json:"-"`
	UpdatedAt  time.Time
}
type Customer struct {
	ID         string      `sql:"primary_key;default 'SE'||nextval('customerIDGenerator')" json:"id"`
	Name       string      `gorm:"size:100;not null" json:"name"`
	CreatedAt  time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	KycID      *uint32     `json:"kyc_id"`
	KycDetails *KYCDetails `json:"kyc,omitempty"`
	AccountID  *uint64     `json:"account_number"`
	Account    *Account    `json:"account,omitempty"`
	Mobile     string      `gorm:"size:100;not null;unique" json:"mobile"`
	Email      string      `gorm:"size:100;not null;unique" json:"email"`
	Password   string      `gorm:"size:100;not null;" json:"-"`
	UpdatedAt  time.Time
}
type KYCDetails struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	UserID    string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	KYCNumber string    `gorm:"size:100;not null;unique" json:"kyc_number"`
	Type      KYCType   `gorm:"size:100;not null" json:"type"`
	UpdatedAt time.Time
}
type Account struct {
	ID        uint64          `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Balance   decimal.Decimal `gorm:"type:decimal(20,8);" json:"balance" `
	UserID    string
	UpdatedAt time.Time
	//Transactions []Transaction   `gorm:"foreignKey:ID;references:ID"`
}
type Transaction struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	AccountID uint32
	Type      TransactionType
	Amount    decimal.Decimal `gorm:"type:decimal(20,8);" json:"amount" `
}

type KYCType string
type TransactionType string

const (
	Debit        = TransactionType("Debit")
	Credit       = TransactionType("Credit")
	Aadhar       = KYCType("Aadhar Card")
	Voter        = KYCType("Voter ID")
	AdminRole    = "admin"
	EmployeeRole = "employee"
	CustomerRole = "customer"
)
