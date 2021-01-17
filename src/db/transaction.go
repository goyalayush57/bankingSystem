package db

import "src/src/models"

type TransactionOperation interface {
	Create(models.Transaction) (models.Transaction, error)
	GetFromID(transactionID uint32) (models.Transaction, error)
	GetTransactions(accountID uint32, size int) ([]models.Transaction, error)
	GetAllTransaction(accountID uint32) ([]models.Transaction, error)
	Update(models.Transaction) error
}

type transaction struct{}

func (t *transaction) Create(txn models.Transaction) (models.Transaction, error) {
	var err error
	err = db.Debug().Model(&models.Transaction{}).Create(&txn).Error
	if err != nil {
		return txn, err
	}
	return txn, nil
}
func (t *transaction) GetFromID(transactionID uint32) (models.Transaction, error) {
	txn := models.Transaction{}
	err := db.Debug().Model(models.Transaction{}).Where("id = ?", transactionID).Take(&txn).Error
	return txn, err
}
func (t *transaction) GetTransactions(accountID uint32, size int) ([]models.Transaction, error) {
	txns := []models.Transaction{}
	err := db.Debug().Model(models.Transaction{}).Where("account_id = ?", accountID).Limit(size).Take(&txns).Error
	return txns, err
}
func (t *transaction) GetAllTransaction(accountID uint32) ([]models.Transaction, error) {
	txns := []models.Transaction{}
	err := db.Debug().Model(&models.Transaction{}).Where("account_id = ?", accountID).Take(&txns).Error
	return txns, err
}

func (t *transaction) Update(txn models.Transaction) error {
	var err error

	err = db.Debug().Model(&models.Account{}).Where("id = ? and account_id = ?", txn.ID, txn, txn.AccountID).Updates(
		models.Transaction{
			Amount: txn.Amount,
			Type:   txn.Type,
		},
	).Error
	if err != nil {
		return err
	}
	return nil
}
