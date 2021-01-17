package db

import "src/src/models"

type AdminOperation interface {
	Get(adminID string) (models.Admin, error)
}
type admin struct{}

func NewAdminOperation() AdminOperation {
	return &admin{}
}

func (a *admin) Get(adminID string) (models.Admin, error) {
	admin := models.Admin{}
	err := db.Debug().Model(models.Admin{}).Where("id = ?", adminID).Take(&admin).Error
	return admin, err
}
