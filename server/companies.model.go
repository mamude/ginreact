package main

import (
	"gorm.io/gorm"
	"time"
)

// Company model
type Company struct {
	ID          int `gorm:"primaryKey"`
	OpeningYear int
	Name        *string `gorm:"uniqueIndex:idx_name,sort:asc,not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
