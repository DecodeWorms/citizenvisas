package types

import (
	"gorm.io/gorm"
)

type Citizen struct {
	gorm.Model
	Name    string
	Gender  string
	Age     int
	Country string
}
