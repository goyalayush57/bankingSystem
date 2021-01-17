package service

import (
	"src/src/db"
	"src/src/models"
	"time"
)

type AccountService interface {
	Save(acc models.Account, userID string) (models.Account, error)
	Get(user_id string) (models.Account, error)
	Update(accountID uint64, user_id string, kyc models.Account) (models.Account, error)
	Delete(accountID uint64, user_id string) error
}

type account struct {
	op db.AccountOperation
}

func NewAccountService() AccountService {
	return &account{
		db.NewAccountService(),
	}
}

//user id is prepopulated
func (a *account) Save(acc models.Account, userID string) (models.Account, error) {
	acc.UserID = userID
	return a.op.Create(acc, userID)
}
func (a *account) Get(user_id string) (models.Account, error) {
	return a.op.Get(user_id)
}
func (a *account) Update(acc_id uint64, user_id string, acc models.Account) (models.Account, error) {
	acc.ID = acc_id
	acc.UserID = user_id
	acc.UpdatedAt = time.Now()
	return a.op.Update(acc)
}
func (a *account) Delete(acc_id uint64, user_id string) error {
	return a.op.Delete(acc_id, user_id)
}
