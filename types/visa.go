package types

import (
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
