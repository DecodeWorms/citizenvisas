package webservices

import (
	"context"

	"gorm.io/gorm"
)

type Citizen struct {
	gorm.Model
	Name    string
	Gender  string
	Age     int
	Country string
}

type CitizenServices interface {
	Create(ctx context.Context, data Citizen, db *gorm.DB) error
	Citizens(ctx context.Context, db *gorm.DB) ([]Citizen, error)
	Citizen(ctx context.Context, db *gorm.DB, id string) (Citizen, error)
	Update(ctx context.Context, db *gorm.DB, data Citizen, id string) error
	Delete(ctx context.Context, db *gorm.DB, data Citizen, id string) error
}
