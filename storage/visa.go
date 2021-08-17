package storage

import (
	"citizenvisas/types"
	"context"

	"gorm.io/gorm"
)

type VisaDataStore struct {
}

type Data struct {
	Data    string
	NewData string
}

func (visa VisaDataStore) Create(ctx context.Context, data types.Visa, db *gorm.DB) error {
	return db.Create(&data).Error
}

func (visa VisaDataStore) Visas(ctx context.Context, db *gorm.DB) ([]types.Visa, error) {
	var visas []types.Visa
	return visas, db.Find((&visas)).Error
}

func (visa VisaDataStore) Visa(ctx context.Context, db *gorm.DB, id string) (types.Visa, error) {
	var result types.Visa

	return result, db.Model(&types.Visa{}).Where("visa_number = ?", id).Find(&result).Error

}

func (visa VisaDataStore) Update(ctx context.Context, db *gorm.DB, data types.Visa, visaNumber Data) error {
	return db.Model(&types.Visa{}).Where("visa_number = ?", visaNumber.Data).Update("visa_number", visaNumber.NewData).Error
}

func (visa VisaDataStore) Delete(ctx context.Context, db *gorm.DB, data types.Visa, id string) error {
	return db.Model(&types.Visa{}).Where("visa_number = ?", id).Delete(&data).Error
}
