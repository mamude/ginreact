package main

import (
	"time"

	"gorm.io/gorm"
)

// Market model
type Market struct {
	ID                 int `gorm:"primaryKey"`
	CategoryBusinessID int
	CategoryBusiness   CategoryBusiness `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name               *string          `gorm:"uniqueIndex:idx_market_name,sort:asc,not null"`
	Description        string
	URL                string
	Email              string
	Phone              string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
