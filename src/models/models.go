package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Admin struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	KycID      *uint32   `json:"kyc_id"`
	KycDetails KYCDetails
	AccountID  *uint64 `json:"account_id"`
	Account    Account
	Mobile     string `gorm:"size:100;not null;unique" json:"mobile"`
	Email      string `gorm:"size:100;not null;unique" json:"email"`
	Password   string `gorm:"size:100;not null;" json:"password"`
}
type Employee struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	KycID      *uint32   `json:"kyc_id"`
	KycDetails KYCDetails
	AccountID  *uint64 `json:"account_id"`
	Account    Account
	Mobile     string `gorm:"size:100;not null;unique" json:"mobile"`
	Email      string `gorm:"size:100;not null;unique" json:"email"`
	Password   string `gorm:"size:100;not null;" json:"password"`
}
type Customer struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	KycID      *uint32   `json:"kyc_id"`
	KycDetails KYCDetails
	AccountID  *uint64 `json:"account_number"`
	Account    Account
	Mobile     string `gorm:"size:100;not null;unique" json:"mobile"`
	Email      string `gorm:"size:100;not null;unique" json:"email"`
	Password   string `gorm:"size:100;not null;" json:"password"`
}
type KYCDetails struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint32
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	KYCNumber string    `gorm:"size:100;not null;unique" json:"mobile"`
	Type      KYCType   `gorm:"size:100;not null" json:"type"`
}
type Account struct {
	ID        uint32          `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Balance   decimal.Decimal `gorm:"type:decimal(20,8);" json:"balance" `
	UserID    *uint32
	//Transactions []Transaction   `gorm:"foreignKey:ID;references:ID"`
}
type Transaction struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	AccountID uint32
	Type      TransactionType
	Amount    decimal.Decimal `gorm:"type:decimal(20,8);" json:"amount" `
}

type KYCType string
type TransactionType string

const (
	Debit  = TransactionType("Debit")
	Credit = TransactionType("Credit")
	Aadhar = KYCType("Aadhar Card")
	Voter  = KYCType("Voter ID")
)
