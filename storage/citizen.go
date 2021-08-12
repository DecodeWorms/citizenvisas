package storage

import (
	"citizenvisas/webservices"
	"context"

	"gorm.io/gorm"
)

type DataStore struct {
}

type ChangeData struct {
	Data    string
	NewData string
}

func (store DataStore) Create(ctx context.Context, data webservices.Citizen, db *gorm.DB) error {
	err = db.Create(&data).Error
	return err
}

func (store DataStore) Citizens(ctx context.Context, db *gorm.DB) ([]webservices.Citizen, error) {

	var res []webservices.Citizen

	return res, db.Find(&res).Error
}

func (store DataStore) Citizen(ctx context.Context, db *gorm.DB, id string) (webservices.Citizen, error) {
	var res webservices.Citizen
	return res, db.Model(&webservices.Citizen{}).Where("name =?", id).Find(&res).Error
}

func (store DataStore) Update(ctx context.Context, db *gorm.DB, data webservices.Citizen, id ChangeData) error {

	return db.Model(&webservices.Citizen{}).Where("name = ?", id.Data).Update("name", id.NewData).Error
}

func (store DataStore) Delete(ctx context.Context, db *gorm.DB, id string, data webservices.Citizen) error {
	return db.Model(&webservices.Citizen{}).Where("name = ?", id).Delete(&data).Error
}
