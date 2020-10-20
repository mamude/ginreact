package main

import (
	"gorm.io/gorm"
	"time"
)

// ShoppingCart model
type ShoppingCart struct {
	ID        int `gorm:"primaryKey"`
	MarketID  int
	Market    Market `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductID int
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Price     float64
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
