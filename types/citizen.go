package types

import "time"

type Citizen struct {
	ID        int64     `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Name      string     `gorm:"name" json:"name,omitempty"`
	Gender    string     `gorm:"gender" json:"gender,omitempty"`
	Age       int        `gorm:"age" json:"age,omitempty"`
	Country   string     `gorm:"country" json:"country,omitempty"`
	CreatedAt time.Time  `gorm:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time  `gorm:"updated_at" json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"deleted_at" json:"deleted_at,omitempty"`
}

