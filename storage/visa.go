package storage

import (
	"citizenvisas/webservices"
	"context"

	"gorm.io/gorm"
)

type VisaDataStore struct {
}

type Data struct {
	Data    string
	NewData string
}

func (visa VisaDataStore) Create(ctx context.Context, data webservices.Visa, db *gorm.DB) error {
	return db.Create(&data).Error
}

func (visa VisaDataStore) Visas(ctx context.Context, db *gorm.DB) ([]webservices.Visa, error) {
	var visas []webservices.Visa
	return visas, db.Find((&visas)).Error
}

func (visa VisaDataStore) Visa(ctx context.Context, db *gorm.DB, id string) (webservices.Visa, error) {
	var result webservices.Visa

	return result, db.Model(&webservices.Visa{}).Where("visa_number = ?", id).Find(&result).Error

}

func (visa VisaDataStore) Update(ctx context.Context, db *gorm.DB, data webservices.Visa, visaNumber Data) error {
	return db.Model(&webservices.Visa{}).Where("visa_number = ?", visaNumber.Data).Update("visa_number", visaNumber.NewData).Error
}

func (visa VisaDataStore) Delete(ctx context.Context, db *gorm.DB, data webservices.Visa, id string) error {
	return db.Model(&webservices.Visa{}).Where("visa_number = ?", id).Delete(&data).Error
}
