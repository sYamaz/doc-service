package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	AutoUUIDModel struct {
		ID        string `gorm:"primaryKey;size:255;default:gen_random_uuid()"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	ManualIDModel struct {
		ID        string `gorm:"primaryKey"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}
)
