package main

import (
	"gorm.io/gorm"
	"time"
)

// Market model
type Market struct {
	ID          int     `gorm:"primaryKey"`
	Name        *string `gorm:"uniqueIndex:idx_market_name,sort:asc,not null"`
	Description string
	URL         string
	Email       string
	Phone       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
