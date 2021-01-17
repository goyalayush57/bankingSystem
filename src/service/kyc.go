package service

import (
	"src/src/db"
	"src/src/models"
	"time"
)

type KycService interface {
	Save(kyc models.KYCDetails, userID string) (models.KYCDetails, error)
	Get(user_id string) (models.KYCDetails, error)
	Update(kycID uint32, user_id string, kyc models.KYCDetails) (models.KYCDetails, error)
	Delete(kycID uint32, user_id string) error
}

func NewKycService() KycService {
	return &kyc{
		db.NewKycOperation(),
	}
}

type kyc struct {
	op db.KycOperation
}

//user id is prepopulated
func (k *kyc) Save(kyc models.KYCDetails, userID string) (models.KYCDetails, error) {
	kyc.UserID = userID
	return k.op.Create(kyc)
}
func (k *kyc) Get(user_id string) (models.KYCDetails, error) {
	return k.op.Get(user_id)
}
func (k *kyc) Update(kycID uint32, user_id string, kyc models.KYCDetails) (models.KYCDetails, error) {
	kyc.ID = kycID
	kyc.UserID = user_id
	kyc.UpdatedAt = time.Now()
	return k.op.Update(kyc)
}
func (k *kyc) Delete(kycID uint32, user_id string) error {
	return k.op.Delete(kycID, user_id)
}
