package types

import (
	"time"
)

type Visa struct {
	ID            string     `gorm:"id" json:"id,omitempty"`
	VisaNumber    string     `gorm:"visa_number" json:"visa_number,omitempty"`
	Country       string     `gorm:"country" json:"country,omitempty"`
	DateCollected string     `gorm:"date_collected" json:"date_collected,omitempty"`
	DateExpired   string     `gorm:"date_expired" json:"date_expired,omitempty"`
	CreatedAt     time.Time  `gorm:"created_at" json:"created_at,omitempty"`
	UpdatedAt     time.Time  `gorm:"updated_at" json:"updated_at,omitempty"`
	DeletedAt     *time.Time `gorm:"deleted_at" json:"deleted_at,omitempty"`
}
