package types

import "gorm.io/gorm"

type Visa struct {
	gorm.Model
	CitizenId     string
	Name          string
	VisaNumber    string
	Country       string
	DateCollected string
	DateExpired   string
}
