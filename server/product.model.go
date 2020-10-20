package main

import (
	"gorm.io/gorm"
	"time"
)

// Product model
type Product struct {
	ID          int `gorm:"primaryKey"`
	MarketID    int
	Market      Market `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string
	Description string
	Price       float64
	Amount      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
