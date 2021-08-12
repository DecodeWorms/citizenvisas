package webservices

import (
	"context"

	"gorm.io/gorm"
)

type Visa struct {
	gorm.Model
	CitizenId     int
	VisaNumber    string
	Country       string
	DateCollected string
	DateExpired   string
}

type VisaServices interface {
	Create(ctx context.Context, data Visa, db *gorm.DB) error
	Citizens(ctx context.Context, db *gorm.DB) ([]Visa, error)
	Citizen(ctx context.Context, db *gorm.DB, id string) (Visa, error)
	Update(ctx context.Context, db *gorm.DB, data Visa, id string) error
	Delete(ctx context.Context, db *gorm.DB, data Visa, id string) error
}
