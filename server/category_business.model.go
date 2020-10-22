package main

import (
	"time"

	"gorm.io/gorm"
)

// CategoryBusiness model
type CategoryBusiness struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
