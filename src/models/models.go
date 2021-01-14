package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Admin struct {
	ID         uint32     `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	KycID      uint32     `gorm:default:1 json:"kyc_id"`
	KycDetails KYCDetails `gorm:"foreignkey:ID;references:kyc_id"`
	AccountID  uint64     `json:"account_id"`
	Account    Account    `gorm:"foreignkey:ID;references:AccountID"`
	Mobile     string     `gorm:"size:100;not null;unique" json:"mobile"`
	Email      string     `gorm:"size:100;not null;unique" json:"email"`
	Password   string     `gorm:"size:100;not null;" json:"password"`
}
type Employee struct {
	ID         uint32     `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	KycID      uint32     `json:"kyc_id"`
	KYCDetails KYCDetails `gorm:"foreignkey:ID;references:KycID"`
	AccountID  uint64     `json:"account_id"`
	Account    Account    `gorm:"foreignkey:ID;references:AccountID"`
	Mobile     string     `gorm:"size:100;not null;unique" json:"mobile"`
	Email      string     `gorm:"size:100;not null;unique" json:"email"`
	Password   string     `gorm:"size:100;not null;" json:"password"`
}
type Customer struct {
	ID         uint32     `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	KycID      uint32     `gorm:default:1 json:"kyc_id"`
	KycDetails KYCDetails `gorm:"foreignkey:ID;references:KycID"`
	AccountID  uint64     `json:"account_id"`
	Account    Account    `gorm:"foreignkey:ID;references:AccountID"`
	Mobile     string     `gorm:"size:100;not null;unique" json:"mobile"`
	Email      string     `gorm:"size:100;not null;unique" json:"email"`
	Password   string     `gorm:"size:100;not null;" json:"password"`
}
type KYCDetails struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	KYCNumber string    `gorm:"size:100;not null;unique" json:"mobile"`
	Type      KYCType   `gorm:"size:100;not null" json:"type"`
}
type Account struct {
	ID           uint32          `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt    time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Balance      decimal.Decimal `gorm:"type:decimal(20,8);" json:"balance" `
	Transactions []Transaction   `gorm:"foreignKey:ID;references:ID"`
}
type Transaction struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Type      TransactionType
	Amount    decimal.Decimal `gorm:"type:decimal(20,8);" json:"amount" `
}

type KYCType string
type TransactionType string
