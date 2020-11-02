package main

import (
	"gorm.io/gorm"
	"time"
)

// CategoryBusiness model
type CategoryBusiness struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
